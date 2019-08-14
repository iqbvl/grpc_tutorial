package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello gRPC")
	strName := "Alexander Graham Bell"
	int32Age := int32(100)

	elliot := &Person{
		Name: &strName,
		Age:  &int32Age,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Error Marshall", err)
	}

	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("Error Unmarhsall", err)
	}

	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
}
