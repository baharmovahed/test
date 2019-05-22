package main

import (
	"fmt"

	emitter "github.com/emitter-io/go/v2"
)

const channelKey = "bAfzAwn_Afbcvv6-QVhBWkZotan6EbKK"

func main() {

	c, _ := emitter.Connect("tcp://198.143.182.157:1883", func(_ *emitter.Client, msg emitter.Message) {
		// fmt.Printf("[emitter] -> [A] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
		topic , shit := msg.Topic() , "stats/52540026181c/"
		if ( topic != shit ){
			fmt.Printf(topic)
		}
	})

	fmt.Println("[emitter] <- [A] subscribing to 'sdk-integration-test/...'")
	c.Subscribe(channelKey, "+/", nil)
	for {

	}
}
