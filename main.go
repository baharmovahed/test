package main

import (
	"fmt"
	"encoding/json"

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	

	emitter "github.com/emitter-io/go/v2"
)

const channelKey = "bAfzAwn_Afbcvv6-QVhBWkZotan6EbKK"


type Test3 struct {
  gorm.Model
  Topic string
  Payload []byte
}



type User struct {
  gorm.Model
  UserName string
  UserID string
  Phone string
  Mail string
}

// type Devices struct {
//   gorm.Model
//   DevName string
//   DevID string
// }

// type Temp struct {
// 	gorm.Model
// 	ParentID string
// 	Min []byte
// 	Max []byte
// }

// type Lumen struct {
// 	gorm.Model
// 	ParentID string
// 	Min []byte
// 	Max []byte
// }

// type Humid struct {
// 	gorm.Model
// 	ParentID string
// 	Min []byte
// 	Max []byte
// }

// type Soil struct {
// 	gorm.Model
// 	ParentID string
// 	SoilOneMin []byte
// 	SoilOneMax []byte
// 	SoilTwoMin []byte
// 	SoilTwoMax []byte
// 	SoilThreeMin []byte
// 	SoilThreeMax []byte
// 	SoilFourMin []byte
// 	SoilFourMax []byte
// }


func main() {

	//conect to DB
	db, err := gorm.Open("mysql", "root:@/iotree?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//connect to broker
	c, _ := emitter.Connect("tcp://198.143.182.157:1883", func(_ *emitter.Client, msg emitter.Message) {

		topic , trash  := msg.Topic() , "stats/52540026181c/" 

		if ( topic != trash ){

			fmt.Println(topic)
			// fmt.Println(msg.Payload())

			var dat map[string]interface{}
			if err := json.Unmarshal(msg.Payload(), &dat); err != nil {
        panic(err)
			}
			
			fmt.Println(dat)

			db.AutoMigrate(&Test3{})
			db.Create(&Test3{Topic: topic , Payload: msg.Payload() })
			
			
		}
	})

	fmt.Println("[emitter] <- [A] subscribing to 'sdk-integration-test/...'")
	c.Subscribe(channelKey, "+/", nil)
	for {

	}
}

