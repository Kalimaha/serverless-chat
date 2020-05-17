package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	guuid "github.com/google/uuid"
	"time"
)

type Message struct {
	MessageId string
	ChatId    string
	CreatedAt string
	Text      string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

	message := Message{
		MessageId: guuid.New().String(),
		ChatId:    "42",
		CreatedAt: time.Now().Format(time.RFC3339),
		Text:      "Hallo, world!",
	}

	av, err := dynamodbattribute.MarshalMap(message)
	if err != nil {
		errorMessage := fmt.Sprintf("Error in MarshalMap: %s", err.Error())
		return events.APIGatewayProxyResponse{Body: errorMessage, StatusCode: 422}, nil
	}

	tableName := "geobricks-serverless-chat-db-test"
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		errorMessage := fmt.Sprintf("Error in PutItem: %s", err.Error())
		return events.APIGatewayProxyResponse{Body: errorMessage, StatusCode: 422}, nil
	}

	body, err := json.Marshal(message)
	if err != nil {
		errorMessage := fmt.Sprintf("Error in Marshal: %s", err.Error())
		return events.APIGatewayProxyResponse{Body: errorMessage, StatusCode: 422}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
