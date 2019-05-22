package main

import (
	"fmt"
	"os"

	emitter "github.com/emitter-io/go/v2"
)

const channelKey = "-VxE2t5me8RyGRxbzqdvH_mcnMkYMMJC"

func main() {

	c, _ := emitter.Connect("tcp://198.143.182.157:1883", func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("[emitter] -> [A] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})

	fmt.Println("[emitter] <- [B] publishing to 'sdk-integration-test/'")
	c.Publish(channelKey, "test4/", os.Args[1])
}
