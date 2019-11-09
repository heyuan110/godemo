package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
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

func main() {
	http.Handle("/",websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234",nil);err != nil {
		log.Fatal("ListenAndServer:",err)
	}


}