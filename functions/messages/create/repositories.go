package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func SaveMessage(message Message) (ok bool, error error) {
	db := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

	av, err := dynamodbattribute.MarshalMap(message)
	if err != nil {
		//errorMessage := fmt.Sprintf("Error in MarshalMap: %s", err.Error())
		//return events.APIGatewayProxyResponse{Body: errorMessage, StatusCode: 422}, nil
		return false, err
	}

	tableName := "geobricks-serverless-chat-db-test"
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		//errorMessage := fmt.Sprintf("Error in PutItem: %s", err.Error())
		//return events.APIGatewayProxyResponse{Body: errorMessage, StatusCode: 422}, nil
		return false, err
	}

	return true, nil
}
