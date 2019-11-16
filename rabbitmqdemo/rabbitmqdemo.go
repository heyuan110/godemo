package main

import (
	"flag"
	"fmt"
	"godemo/rabbitmq"
)

var (
	mqHost = flag.String("h", "127.0.0.1", "AMQP URI")
	mqPort = flag.Int("p",5672,"Port")
	mqUserName = flag.String("u","guest","User Name")
	mqPassword = flag.String("pwd","guest","Password")
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

func init() {
	flag.Parse()
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
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",*mqUserName,*mqPassword,*mqHost,*mqPort)
	mq := rabbitmq.New(url,ex)
	mq.RegisterProducer(c)
	//mq.RegisterConsumer(c)
	mq.Start()
}
