package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	guuid "github.com/google/uuid"
	"time"
)

func RetrieveMessages(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	filter := expression.Name("PrimaryKey").BeginsWith("MESSAGE_")
	proj := expression.NamesList(expression.Name("PrimaryKey"), expression.Name("SecondaryKey"), expression.Name("Data"), expression.Name("CreatedAt"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	if err != nil {
		returnFormattedError(err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TableName),
	}

	result, err := dbSession().Scan(params)
	if err != nil {
		returnFormattedError(err)
	}

	//item := Message{}
	//for _, i := range result.Items {
	//	err = dynamodbattribute.UnmarshalMap(i, &item)
	//}
	//body, err := json.Marshal(item)
	//if err != nil {
	//	returnFormattedError(err)
	//}

	out := make([]Message, len(result.Items))
	for i, v := range result.Items {
		tmp := Message{}
		dynamodbattribute.UnmarshalMap(v, &tmp)
		out[i] = tmp
	}

	body, err := json.Marshal(out)
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

func buildMessage() (message Message) {
	return Message{
		PrimaryKey:		fmt.Sprintf("MESSAGE_%s", guuid.New().String()),
		SecondaryKey:	fmt.Sprintf("CHAT_%s", "42"),
		CreatedAt: 		time.Now().Format(time.RFC3339),
		Data:      		"Hallo, world!",
	}
}
