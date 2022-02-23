package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	ps "github.com/kotakanbe/go-pingscanner"
	"golang.org/x/crypto/ssh"
)

var (
	cidr        string
	username    string
	password    string
	netconfPort string
)

func main() {
	getAndParseFlags()
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
		fmt.Println(err)
	} else {
		if len(aliveIPs) < 1 {
			fmt.Println("No alive hosts")
		}
		for _, ip := range aliveIPs {
			d, err := getNetconfDataForHost(ip + netconfPort)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(d.IP.IPv4.Address[0].IP)
			}
		}
	}
}

func getAndParseFlags() {
	flag.StringVar(&cidr, "cidr", "", "CIDR block to scan")
	flag.StringVar(&username, "username", "admin", "The SSH username")
	flag.StringVar(&password, "password", "", "The SSH password")
	flag.StringVar(&netconfPort, "port", ":22", "The NETCONF port to use")
	flag.Parse()
}

func getNetconfDataForHost(host string) (Data, error) {
	fmt.Println("Trying", host)
	var returnData Data
	sshConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 2,
	}

	session, err := netconf.DialSSH(host, sshConfig)

	if err != nil {
		return returnData, err
	}
	defer session.Close()

	reply, err := session.Exec(netconf.MethodGet("xpath", ""))

	err = xml.Unmarshal([]byte(reply.Data), &returnData)

	if err != nil {
		return returnData, err
	}

	return returnData, nil
}
