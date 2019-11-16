package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

//var mqConn *amqp.Connection
//var mqChan *amqp.Channel

type Producer interface {
	Message() string
}

type Consumer interface {
	ConsumeMessage([]byte) error
}

type RabbitMQ struct {
	serverURL string
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

func New(serverUrl string,ex *Exchange) *RabbitMQ  {
	return &RabbitMQ{
		serverURL: 	  serverUrl,
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

func (r *RabbitMQ)mqConnect() error{
	var err error
	logInfo("connect to ",r.serverURL)
	r.connection,err = amqp.Dial(r.serverURL)
	if err != nil{
		logError(err,"Failed to connect Rabbitmq server ")
	}else{
		logInfo("connected successful!")
	}
	return err
}

func (r *RabbitMQ)mqChannel()  error{
	var err error
	r.channel,err= r.connection.Channel()
	if err != nil{
		logError(err,"Failed to open a channel")
	}else{
		logInfo("open a channel")
	}
	return err
}

func (r *RabbitMQ)exchangeDeclare() error {
	err :=  r.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		logError(err,"Failed to declare exchange")
	}else{
		logInfo("exchange declare:",r.exchangeName,r.exchangeType)
	}
	return err
}

func (r *RabbitMQ)queueDeclare() error {
	_,err := r.channel.QueueDeclare(r.queueName,true,false,false,true,nil)
	if err != nil{
		logError(err,"Failed to declare the queue")
	}else{
		logInfo("queue declare:",r.queueName)
	}
	return err
}

func (r *RabbitMQ)queueBind() error {
	err := r.channel.QueueBind(r.queueName,r.routingKey,r.exchangeName,true,nil)
	if err != nil {
		logError(err,"Failed to bind the queue")
	}else{
		logInfo("queue bind:","queue->",r.queueName,",routing key->",r.routingKey,"exchange->",r.exchangeName)
	}
	return err
}

func (r *RabbitMQ)publishMessage(msg string) error {
	err := r.channel.Publish(r.exchangeName,r.routingKey,false,false,amqp.Publishing{
		ContentType: "text/plain",
		Timestamp:       time.Time{},
		Body:            []byte(msg),
	})
	if err != nil{
		logError(err,"Failed to publish the message")
	}else{
		logInfo("publish to exchange:",r.exchangeName," message:",msg)
	}
	return err
}

func (r *RabbitMQ)consumeMessage(consumer Consumer) {
	err :=r.channel.Qos(1,0,true)
	msgList,err := r.channel.Consume(r.queueName,"test",false,false,false,false,nil)
	if err != nil{
		logError(err,"Failed to consume message")
		return
	}
	for msg :=  range msgList {
		err := consumer.ConsumeMessage(msg.Body)
		if err != nil {
			logError(err,"Consumption Message is failure")
			err = msg.Ack(true)
			if err != nil {
				logError(err,"Consumption message failed, call ack has an exception,messageID: "+msg.MessageId)
			}
		}else{
			//multiple value must false
			err = msg.Ack(false)
			if err != nil {
				logError(err,"Consumption message succeed, call ack has an exception,messageID: "+msg.MessageId)
			}
		}
	}
}

func (r *RabbitMQ)assembleQueue() error {
	//declare a exchange
	err := r.exchangeDeclare()
	if err != nil {
		return err
	}
	//declare a queue
	err = r.queueDeclare()
	if err != nil {
		return err
	}
	//bind exchange and queue
	err = r.queueBind()
	if err != nil {
		return err
	}
	return  err
}

func (r *RabbitMQ)listenProducer(producer Producer)  {
	if r.mqConnect() != nil{
		return
	}else{
		defer r.mqConnectClose()
	}
	if r.mqChannel() != nil{
		return
	}else{
		defer r.mqChannelClose()
	}
	if r.assembleQueue() != nil {
		return
	}
	//publish message
	if r.publishMessage(producer.Message()) != nil {
		return
	}
}

func (r *RabbitMQ)listerConsumer(consumer Consumer)  {
	if r.mqConnect() != nil{
		return
	}else{
		defer r.mqConnectClose()
	}
	if r.mqChannel() != nil{
		return
	}else{
		defer r.mqChannelClose()
	}
	if r.assembleQueue() != nil {
		return
	}
	r.consumeMessage(consumer)
}

func (r *RabbitMQ)mqChannelClose() {
	err := r.channel.Close()
	if err != nil{
		logError(err,"Failed to close channel")
	}else{
		logInfo("channel closed")
	}
}

func (r *RabbitMQ)mqConnectClose() {
	err := r.connection.Close()
	if err != nil{
		logError(err,"Failed to close connection")
	}else{
		logInfo("connection closed")
	}
}

func (r *RabbitMQ)mqClose()  {
	r.mqChannelClose()
	r.mqConnectClose()
}

func logInfo(a ...interface{})  {
	a = append(a[:1],a[0:]...)
	a[0] = time.Now().String()
	fmt.Println(a...)
}

func logError(err error,msg string)  {
	if err != nil {
		fmt.Printf("%s %s:%s\n",time.Now().String(),msg,err)
	}
}
