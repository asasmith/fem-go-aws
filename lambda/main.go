package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// take in a payload and do somethign with it
func HandleRequest(event MyEvent) (string, error) { 
	if event.Username == "" {
		return "", fmt.Errorf("Username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s!", event.Username), nil
}

func main() {
	lambda.Start(HandleRequest)
}
