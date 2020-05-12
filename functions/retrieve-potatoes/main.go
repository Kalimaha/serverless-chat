package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	RequestId float64 `json:"requestId"`
}

type Response struct {
	ResponseId float64 `json:"responseId"`
	Message    string  `json:"message"`
}

func Handler(request Request) (Response, error) {
	return Response{
		ResponseId: request.RequestId,
		Message:    fmt.Sprintf("Hallo, world!"),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
