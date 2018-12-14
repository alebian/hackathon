package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
  return fmt.Sprintf("Hello %s!", event.Name ), nil
}

func main() {
  lambda.Start(HandleRequest)
}
