package main

import (
	"fmt"
	"github.com/quinont/yeelight"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var y *yeelight.Yeelight

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	y.ThrowAlarm()
}

func main() {
	// Tomando la ip del foco.
	iplamp := os.Getenv("IP_LAMP")
	if iplamp == "" {
		log.Fatal("No se seteo el IP_LAMP")
	}

	y = yeelight.New(iplamp + ":55443")

	on, err := y.GetProp("power")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Power es %s", on[0].(string))
	if on[0].(string) == "off" {
		fmt.Printf(" Prendiendo la lampara.")
		y.SetPower(true)
	}
	fmt.Printf("\n")

	// Tomando la ip del Server Mosquitto.
	ipmqtt := os.Getenv("IP_MQTT")
	if iplamp == "" {
		log.Fatal("No se seteo el IP_LAMP")
	}
	server := "tcp://" + ipmqtt + ":1883"
	topic := os.Getenv("TOPIC")
	if iplamp == "" {
		log.Fatal("No se seteo el TOPIC")
	}
	qos := 0
	hostname, _ := os.Hostname()
	clientid := hostname + strconv.Itoa(time.Now().Second())
	connOpts := MQTT.NewClientOptions().AddBroker(server).SetClientID(clientid).SetCleanSession(true)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}
	<-c
}
