package main

import (
	"encoding/json"
	"fmt"
	kafka "github.com/RedHatInsights/strimzi-client-go/apis/kafka.strimzi.io/v1beta1"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("kafka.json")

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	k := kafka.KafkaList{}
	err = json.Unmarshal(data, &k)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
