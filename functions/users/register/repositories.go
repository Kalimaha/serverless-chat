package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func UpdateUser(dbSession dynamodbiface.DynamoDBAPI, user User) (ok bool, error error) {
	userItem, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return false, err
	}

	input := &dynamodb.PutItemInput{
		Item:      userItem,
		TableName: aws.String(TableName),
	}

	_, err = dbSession.PutItem(input)
	if err != nil {
		return false, err
	}

	return true, nil
}
