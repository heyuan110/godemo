package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"flag"
	"fmt"
	"godemo/rabbitmq"
)

var (
	mqHost = flag.String("h", "127.0.0.1", "AMQP URI")
	mqPort = flag.Int("p",5672,"Port")
	mqUserName = flag.String("u","guest","User Name")
	mqPassword = flag.String("pwd","guest","Password")
	mqVHost = flag.String("vhost","/","Virtual Hosts")
)

type Chat struct {
	msg string
}

func (c *Chat)Message() string {
	return c.msg
}

func (c *Chat)ConsumeMessage(msgByte []byte) error {
	msg := string(msgByte)
	fmt.Println("I am a consumer, get message :",msg)
	db, err := sql.Open("mysql","root:root@tcp(localhost:3306)/test?charset=utf8")
	if err!=nil{
		rabbitmq.LogError(err,"Failed to connect to database")
		return err
	}
	stmt,err := db.Prepare("INSERT chat_records SET msg=?")
	if err!=nil{
		rabbitmq.LogError(err,"Failed to Prepare to insert")
		return err
	}
	res, err := stmt.Exec(msg)
	if err!=nil{
		rabbitmq.LogError(err,"Failed to insert data to table")
		return err
	}
	id,err :=res.LastInsertId()
	rabbitmq.LogInfo("Insert record id is ",id)
	return  err
}

func init() {
	flag.Parse()
}

func main() {
	msg := "this is a test task"
	c := &Chat{msg:msg}
	serverUrl := rabbitmq.RabbitMQServer{
		Host:     *mqHost,
		Port:     *mqPort,
		User:     *mqUserName,
		Password: *mqPassword,
		VHost: *mqVHost,
	}
	ex:=&rabbitmq.Exchange{
		QuName: "test.rabbitmq",
		RtKey:  "test.rabbitmq.routingkey",
		ExName: "test.rabbitmq.exchange",
		ExType: "direct",
	}
	mq := rabbitmq.New(serverUrl,ex)
	mq.RegisterProducer(c)
	mq.RegisterConsumer(c)
	mq.Start()
}
