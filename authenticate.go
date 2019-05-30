package main

import (
	// "encoding/json"
	"fmt"

	emitter "github.com/emitter-io/go/v2"
)

const channelKey = "KBUaNMCiMKjifJVGJdGKvWCIY_C4qXyA/"

type response1 struct {
    Page   int
    Fruits []string
}

func main() {

	c, _ := emitter.Connect("tcp://198.143.182.157:1883", func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("[emitter] -> [A] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})

	newChannelKey := map[string]string{"key":"_sTGZ2hepotGgdZO9c8MQ8PBb2TPZVhj","channel":"article1"}

	fmt.Println("[emitter] <- [B] publishing to 'sdk-integration-test/'")
	c.Publish(channelKey, "emitter/keygen/", newChannelKey)
}
