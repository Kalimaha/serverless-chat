package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	guuid "github.com/google/uuid"
	"time"
)

func CreateMessage(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := buildMessage()

	_, err := SaveMessage(message)
	if err != nil {
		returnFormattedError(err)
	}

	body, err := json.Marshal(message)
	if err != nil {
		returnFormattedError(err)
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func returnFormattedError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 422}, nil
}

func buildMessage() (message Message) {
	return Message{
		MessageId: guuid.New().String(),
		ChatId:    "42",
		CreatedAt: time.Now().Format(time.RFC3339),
		Text:      "Hallo, world!",
	}
}
