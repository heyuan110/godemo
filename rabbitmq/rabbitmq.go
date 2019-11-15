package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var mqConn *amqp.Connection
var mqChan *amqp.Channel

type Producer interface {
	Message() string
}

type Consumer interface {
	ConsumeMessage([]byte) error
}

type RabbitMQ struct {
	connection *amqp.Connection
	channel *amqp.Channel
	queueName   string            // 队列名称
	routingKey  string            // key名称
	exchangeName string           // 交换机名称
	exchangeType string           // 交换机类型
	producerList []Producer
	consumerList []Consumer
	mu sync.RWMutex
}

type Exchange struct {
	QuName  string           // 队列名称
	RtKey   string           // key值
	ExName  string           // 交换机名称
	ExType  string           // 交换机类型
}

func New(ex *Exchange) *RabbitMQ  {
	return &RabbitMQ{
		queueName:    ex.QuName,
		routingKey:   ex.RtKey,
		exchangeName: ex.ExName,
		exchangeType: ex.ExType,
	}
}

func (r *RabbitMQ)Start() {
	for _,producer := range r.producerList {
		go r.listenProducer(producer)
	}
	for _,consumer := range r.consumerList {
		go r.listerConsumer(consumer)
	}
	//sleep 1 second
	time.Sleep(1*time.Second)
}

func (r *RabbitMQ)RegisterProducer(producer Producer)  {
 	r.producerList = append(r.producerList,producer)
}

func (r *RabbitMQ)RegisterConsumer(consumer Consumer)  {
	r.mu.Lock()
	r.consumerList = append(r.consumerList,consumer)
	r.mu.Unlock()
}

func (r *RabbitMQ)mqConnect()  {
	var err error
	userName := "bruce1"
	password := "patpat"
	host := "192.168.8.131"
	port := 5678
	rabbitmqServerUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/",userName,password,host,port)
	mqConn,err = amqp.Dial(rabbitmqServerUrl)
	r.connection = mqConn
	showError(err,"Failed to connect Rabbitmq server ")
	mqChan,err = mqConn.Channel()
	r.channel = mqChan
	showError(err,"Failed to open a channel")
}

func (r *RabbitMQ)mqClose()  {
	err := r.channel.Close()
	showError(err,"Failed to close channel")
	err = r.connection.Close()
	showError(err,"Failed to close connection")
}

func (r *RabbitMQ)exchangeDeclare() error {
	err := r.channel.ExchangeDeclarePassive(r.exchangeName,r.exchangeType,true,false,false,true,nil)
	if err != nil {
		err =  r.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, true, nil)
		if err != nil {
			showError(err,"Failed to declare exchange")
		}
	}
	return err
}

func (r *RabbitMQ)queueDeclare() error {
	_,err := r.channel.QueueDeclarePassive(r.queueName,true,false,false,true,nil)
	if err != nil{
		_,err = r.channel.QueueDeclare(r.queueName,true,false,false,true,nil)
		if err != nil{
			showError(err,"Failed to declare the queue")
		}
	}
	return err
}

func (r *RabbitMQ)queueBind() error {
	err := r.channel.QueueBind(r.queueName,r.routingKey,r.exchangeName,true,nil)
	if err != nil {
		showError(err,"Failed to bind the queue")
	}
	return err
}

func (r *RabbitMQ)publishMessage(msg string) error {
	err := r.channel.Publish(r.exchangeName,r.routingKey,false,false,amqp.Publishing{
		ContentType: "text/plain",
		Timestamp:       time.Time{},
		Body:            []byte(msg),
	})
	showError(err,"Failed to publish the message")
	return err
}

func (r *RabbitMQ)consumeMessage(consumer Consumer) {
	_:=r.channel.Qos(1,0,true)
	msgList,err := r.channel.Consume(r.queueName,"test",false,false,false,false,nil)
	if err != nil{
		showError(err,"Failed to consume message")
		return
	}
	for msg :=  range msgList {
		err := consumer.ConsumeMessage(msg.Body)
		if err != nil {
			showError(err,"Consumption Message is failure")
			err = msg.Ack(true)
			if err != nil {
				showError(err,"Consumption message failed, call ack has an exception,messageID: "+msg.MessageId)
			}
		}else{
			//multiple value must false
			err = msg.Ack(false)
			if err != nil {
				showError(err,"Consumption message succeed, call ack has an exception,messageID: "+msg.MessageId)
			}
		}
	}
}

func (r *RabbitMQ)assemble() error {
	var err error = errors.New("failure")
	//create a connection, create a channel
	if r.channel == nil {
		r.mqConnect()
	}
	//declare a exchange
	if r.exchangeDeclare() != nil {
		return err
	}
	//declare a queue
	if r.queueDeclare() != nil {
		return err
	}
	//bind exchange and queue
	if r.queueBind() != nil {
		return err
	}
	return  nil
}

func (r *RabbitMQ)listenProducer(producer Producer)  {
	//close channel, connnection
	defer r.mqClose()
	if r.assemble() != nil {
		return
	}
	//publish message
	if r.publishMessage(producer.Message()) != nil {
		return
	}
}

func (r *RabbitMQ)listerConsumer(consumer Consumer)  {
	defer r.mqClose()
	if r.assemble() != nil {
		return
	}
	r.consumeMessage(consumer)
}

func showError(err error,msg string)  {
	if err != nil {
		fmt.Printf("%s:%s\n",msg,err)
	}
}
