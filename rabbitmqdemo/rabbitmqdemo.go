package main

import (
	"fmt"
	rabbitmq "godemo/rabbitmqlib"
)

type Chat struct {
	msg string
}

func (c *Chat)Message() string {
	return c.msg
}

func (c *Chat)ConsumeMessage(msgByte []byte) error {
	fmt.Println("I am a consumer, get message :",string(msgByte))
	return  nil
}

func main() {
	msg := "this is a test task"
	c := &Chat{msg:msg}
	ex:=&rabbitmq.Exchange{
		QuName: "test.rabbitmq",
		RtKey:  "test.rabbitmq.routingkey",
		ExName: "test.rabbitmq.exchange",
		ExType: "direct",
	}
	mq := rabbitmq.New(ex)
	mq.RegisterProducer(c)
	mq.RegisterConsumer(c)
	mq.Start()
}
