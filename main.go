package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Juniper/go-netconf/netconf"
	"golang.org/x/crypto/ssh"
)

func main() {
	host := os.Args[1]
	sshConfig := &ssh.ClientConfig{
		User:            os.Args[2],
		Auth:            []ssh.AuthMethod{ssh.Password(os.Args[3])},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	session, err := netconf.DialSSH(host, sshConfig)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer session.Close()

	reply, err := session.Exec(netconf.MethodGet("xpath", ""))
	var returnData Data

	err = xml.Unmarshal([]byte(reply.Data), &returnData)

	if err != nil {
		fmt.Println(err)
		return
	}

	k, err := json.MarshalIndent(returnData, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(k))
}
