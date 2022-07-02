package main

import (
	"fmt"
	"log"

	"github.com/lamhai1401/protobuff-ex/model"
	"google.golang.org/protobuf/proto"
)

// type Person struct {
// 	Name string
// 	Age  int
// }

func main() {

	elliot := model.Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &model.SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(&elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// let's go the other way and unmarshal
	// our protocol buffer into an object we can modify
	// and use
	newElliot := &model.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())

}
