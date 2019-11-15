package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func Echo(ws *websocket.Conn)  {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		if reply == "exit" {
			fmt.Println("Socket disconnect!")
			break
		}

		if reply == "1" {
			reply = "咨询业务"
		}else if reply == "2" {
			reply = "查询订单"
		}

		msg := time.Now().String() + " server:  " + reply

		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}


func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func main() {
	localIp := "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if (ipnet.IP.To4() != nil && strings.HasPrefix(ipnet.IP.String(),"192")) {
				fmt.Println("local ip address ",ipnet.IP.String())
				localIp = ipnet.IP.String()
				break
			}
		}
	}
	http.Handle("/",websocket.Handler(Echo))
	addr := localIp+":1234"
	fmt.Println("start websocket server, address ",addr)
	fmt.Println("waiting client connect...")
	if err := http.ListenAndServe(addr,nil);err != nil {
		log.Fatal("ListenAndServer:",err)
	}
}