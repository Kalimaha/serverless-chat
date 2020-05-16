package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	ResponseId string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestId	:= request.PathParameters["requestId"]
	response 	:= Response{ResponseId: requestId}
	body, _ 	:= json.Marshal(response)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
