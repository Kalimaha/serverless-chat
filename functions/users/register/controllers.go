package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"time"
)

func RegisterUser(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := buildUser(request.RequestContext.ConnectionID, request.RequestContext.EventType)

	_, err := UpdateUser(dbSession(), user)
	if err != nil {
		returnFormattedError(err)
	}

	body, err := json.Marshal(user)
	if err != nil {
		returnFormattedError(err)
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 201}, nil
}

func dbSession() (dbSession dynamodbiface.DynamoDBAPI) {
	sess := session.Must(session.NewSession())

	return dynamodb.New(sess)
}

func returnFormattedError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 422}, nil
}

func buildUser(connectionId string, eventType string) (user User) {
	return User{
		PrimaryKey:	fmt.Sprintf("USER_%s", connectionId),
		Data:    	eventType == "CONNECT",
		CreatedAt: 	time.Now().Format(time.RFC3339),
	}
}
