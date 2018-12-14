package main

import (
	"encoding/json"
	awsEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Parameters struct {
	name string `json:"name"`
}

func Handle(request awsEvents.APIGatewayProxyRequest) (response awsEvents.APIGatewayProxyResponse, err error) {
	jsonBody := make(map[string]string)
	jsonBody["event"] = request.QueryStringParameters["name"]

	body, err := json.Marshal(jsonBody)
	if err != nil {
		return
	}

	response = awsEvents.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}

	return
}

func main() {
	lambda.Start(Handle)
}
