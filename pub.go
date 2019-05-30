package main

import (
	"encoding/json"
	"fmt"
	// "os"

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

	res2D := &response1{
		Page : 2,
		Fruits : []string{"hh","llll","sss"}}
    res2B, _ := json.Marshal(res2D)

	fmt.Println("[emitter] <- [B] publishing to 'sdk-integration-test/'")
	c.Publish(channelKey, "dev/s/", res2B)
}


//os.Args[1]
