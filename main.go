package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
    FirstName string json:"firstName"
    LastName  string json:"lastName"
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    return fmt.Sprintf("Hello %s %s!", name.FirstName, name.LastName), nil
}

func main() {
    lambda.Start(HandleRequest)
}