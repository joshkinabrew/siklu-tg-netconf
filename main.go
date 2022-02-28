package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	ps "github.com/kotakanbe/go-pingscanner"
	"github.com/potato2003/actioncable-client-go"
	"golang.org/x/crypto/ssh"
)

var (
	cidr                    string
	username                string
	password                string
	netconfPort             string
	websocketServerAddress  string
	useWSS                  bool
	actionCableSubscription *actioncable.Subscription
	serverToken             string
	channelName             string
	debug                   bool
	refreshRate             int
	excludedIPs             []string
	excludedIPList          string
)

// ChannelSubscriptionEventHandler ...
type ChannelSubscriptionEventHandler struct {
	actioncable.SubscriptionEventHandler
}

// OnConnected ...
func (h *ChannelSubscriptionEventHandler) OnConnected(se *actioncable.SubscriptionEvent) {
	log.Println("Connected to websocket server", websocketServerAddress)
}

// OnDisconnected ...
func (h *ChannelSubscriptionEventHandler) OnDisconnected(se *actioncable.SubscriptionEvent) {
	fmt.Println("On disconnect")
}

func main() {
	getAndParseFlags()
	connectToWebsocketServer()

	for {
		<-time.After(time.Duration(refreshRate) * time.Second)

		err := scanAndSend()
		if err != nil {
			debugLog(err)
		}
	}
}

func scanAndSend() error {
	scanner := ps.PingScanner{
		CIDR: cidr,
		PingOptions: []string{
			"-c1",
			"-W1",
		},
		NumOfConcurrency: 100,
	}

	aliveIPs, err := scanner.Scan()

	if err != nil {
		return err
	} else {
		if len(aliveIPs) < 1 {
			fmt.Println("No alive hosts")
		}
		debugLog("Found", len(aliveIPs), "hosts")
		for _, ip := range aliveIPs {
			if contains(excludedIPs, ip) {
				continue
			}

			d, err := getNetconfDataForHost(ip + netconfPort)
			if err != nil {
				debugLog(err)
				continue
			} else {
				go sendDataToWebsocketServer(d)
			}
		}
	}

	return nil
}

func getAndParseFlags() {
	flag.StringVar(&cidr, "cidr", "", "CIDR block to scan")
	flag.StringVar(&username, "username", "admin", "The SSH username")
	flag.StringVar(&password, "password", "", "The SSH password")
	flag.StringVar(&netconfPort, "port", ":22", "The NETCONF port to use")
	flag.StringVar(&websocketServerAddress, "server", "", "The websocket server address")
	flag.StringVar(&serverToken, "token", "", "Websocket server password")
	flag.StringVar(&channelName, "channel", "", "Websocket server channel")
	flag.IntVar(&refreshRate, "refresh", 30, "Data fetch rate (in seconds)")
	flag.BoolVar(&useWSS, "secure", true, "Use 'wss:// to connect to the websocket server")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.StringVar(&excludedIPList, "excluded-ips", "", "Excluded IPs from scan (comma separated) e.g. -excluded-ips=1.1.1.1,2.2.2.2")
	flag.Parse()

	excludedIPs = strings.Split(excludedIPList, ",")
}

func getNetconfDataForHost(host string) (Data, error) {
	debugLog("Trying", host)
	var returnData Data
	sshConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 3,
	}

	session, err := netconf.DialSSH(host, sshConfig)

	if err != nil {
		return returnData, fmt.Errorf("%s - %s", host, err)
	}
	defer session.Close()

	reply, err := session.Exec(netconf.MethodGet("xpath", ""))

	err = xml.Unmarshal([]byte(reply.Data), &returnData)

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}

func connectToWebsocketServer() (err error) {
	var scheme string
	if useWSS {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	u, _ := url.Parse(fmt.Sprintf("%s://%s/cable?token=%s", scheme, websocketServerAddress, serverToken))
	log.Println("Connecting to websocket server", u)

	header := http.Header{}
	header.Set("Origin", "127.0.0.1")

	opt := actioncable.NewConsumerOptions()
	opt.SetHeader(&header)

	consumer, err := actioncable.CreateConsumer(u, opt)
	if err != nil {
		log.Println(err)
	}
	consumer.Connect()

	params := map[string]interface{}{
		"token": serverToken,
	}

	id := actioncable.NewChannelIdentifier(channelName, params)
	actionCableSubscription, err = consumer.Subscriptions.Create(id)
	if err != nil {
		log.Println(err)
	}
	actionCableSubscription.SetHandler(&ChannelSubscriptionEventHandler{})

	return err
}

func sendDataToWebsocketServer(d Data) {
	j, _ := json.Marshal(d)
	data := map[string]interface{}{
		"device_data": string(j),
		"time":        time.Now().UnixNano() / int64(time.Millisecond),
	}
	actionCableSubscription.Perform("device_update", data)
}

func debugLog(a ...interface{}) {
	if debug {
		log.Println(a...)
	}
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
