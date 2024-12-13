package exmq

import (
	"fmt"
	"gin-api/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MQTTOptions struct {
	ClientID              string
	Username              string
	Password              string
	KeepAlive             uint
	MessageHandler        func(client mqtt.Client, msg mqtt.Message)
	ConnectHandler        func(client mqtt.Client)
	ConnectionLostHandler func(client mqtt.Client, err error)
	PublishHandler        func(client mqtt.Client, msg mqtt.Message)
	DisconnectHandler     func(client mqtt.Client)
	ErrorHandler          func(client mqtt.Client, err error)
}

func ConnectMQTT(
	options *MQTTOptions,
) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", global.CONFIG.MQTT.Broker, global.CONFIG.MQTT.Port))

	opts.SetClientID(options.ClientID)
	opts.SetUsername(options.Username)
	opts.SetPassword(options.Password)

	if options.MessageHandler != nil {
		opts.SetDefaultPublishHandler(options.MessageHandler)
	}
	if options.ConnectHandler != nil {
		opts.SetOnConnectHandler(options.ConnectHandler)
	}
	if options.ConnectionLostHandler != nil {
		opts.SetConnectionLostHandler(options.ConnectionLostHandler)
	}
	if options.KeepAlive != 0 {
		duration := time.Duration(options.KeepAlive) * time.Hour
		opts.SetKeepAlive(duration)
	} else {
		opts.SetKeepAlive(60)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
