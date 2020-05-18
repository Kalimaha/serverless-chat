package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	guuid "github.com/google/uuid"
	"time"
)

func RegisterUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := buildUser()

	_, err := UpdateUser(dbSession(), message)
	if err != nil {
		returnFormattedError(err)
	}

	body, err := json.Marshal(message)
	if err != nil {
		returnFormattedError(err)
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func dbSession() (dbSession dynamodbiface.DynamoDBAPI) {
	sess := session.Must(session.NewSession())

	return dynamodb.New(sess)
}

func returnFormattedError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 422}, nil
}

func buildUser() (user User) {
	return User{
		PrimaryKey:	guuid.New().String(),
		Active:    	true,
		CreatedAt: 	time.Now().Format(time.RFC3339),
	}
}
