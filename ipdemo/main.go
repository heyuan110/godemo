package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	seelog "github.com/cihub/seelog"
	"net"
	"os"
)

func main() {

	defer seelog.Flush()
	seelog.Info("Hello from Seelog!")

	return

	log.WithFields(log.Fields{
		"animal": "walrus",
		"testkey":"testvalue",
	}).Info("A walrus appears")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//fmt.Println(os.Args)
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	fmt.Println(name)
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
