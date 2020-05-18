package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func SaveMessage(dbSession dynamodbiface.DynamoDBAPI, message Message) (ok bool, error error) {
	messageItem, err := dynamodbattribute.MarshalMap(message)
	if err != nil {
		return false, err
	}

	input := &dynamodb.PutItemInput{
		Item:      messageItem,
		TableName: aws.String(TableName),
	}

	_, err = dbSession.PutItem(input)
	if err != nil {
		return false, err
	}

	return true, nil
}
