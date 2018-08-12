package main

import (
	"encoding/json"
	"fmt"
	"github.com/serinth/serverless-emailer/api"
)

func local() {

	var req api.SendEmailRequest

	if err := json.Unmarshal([]byte(`{ "to": "toField", "cc": [ "one", "two"] }`), &req); err != nil {
		fmt.Println("Error: %v", err)
		panic("Failed to unmarshal with error")
	}

	fmt.Println("Decoded is: %v", *req.To)

	for _, v := range req.CC {
		fmt.Println("Decoded is: %v", *v)
	}

}
