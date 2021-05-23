package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsServer = "nats://nats:4222"
)

var(
	nc *nats.Conn
	sender string
	listenFor string
	sendMessage bool
)

func init() {
	sender = os.Getenv("NAME")
	listenFor = os.Getenv("LISTEN_FOR")
	sendMessage = (os.Getenv("SEND") == "true")

	// Connect to a server
	var err error
  nc, err = nats.Connect(natsServer)
	if err != nil {
		panic(fmt.Errorf("error connecting to nats: %s", err))
	}
	fmt.Println("Connected to NATS server.")
	
	nc.Subscribe("message:" + listenFor, messageReceiver)
	fmt.Println("Listening for messages from", listenFor)
}

func messageReceiver(m *nats.Msg) {
	fmt.Printf("Sender %s slept for: %s\n", listenFor, string(m.Data))
}

func main() {
	defer nc.Close()

	for {
		if sendMessage {
			rndTime := int64(rand.Intn(3000))
	
			fmt.Println("Sleeping for", rndTime, "ms")
			time.Sleep(time.Duration(rndTime)* time.Millisecond)
			nc.Publish("message:" + sender, []byte(fmt.Sprintf("%d ms", rndTime)))
		}
	}
}
