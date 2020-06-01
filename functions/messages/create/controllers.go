package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	guuid "github.com/google/uuid"
	"time"
)

type LambdaRequest struct {
	Data string `json:"data"`
}

type Pippo struct {
	PrimaryKey string
	Data bool
}

func CreateMessage(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaRequest := LambdaRequest{}
	json.Unmarshal([]byte(request.Body), &lambdaRequest)

	message := buildMessage(lambdaRequest.Data)

	_, err := SaveMessage(dbSession(), message)
	if err != nil {
		returnFormattedError(err)
	}

	//body, err := json.Marshal(message)
	//if err != nil {
	//	returnFormattedError(err)
	//}

	// Retrieve active users
	conditionOne := expression.Name("PrimaryKey").BeginsWith("USER_")
	conditionTwo := expression.Name("Data").Equal(expression.Value(true))
	//expr, err := expression.NewBuilder().WithCondition(conditionOne.And(conditionTwo)).Build()

	//expr, err := expression.NewBuilder().WithCondition(conditionOne).Build()

	//filt := expression.Name("PrimaryKey").BeginsWith("USER_")
	proj := expression.NamesList(expression.Name("PrimaryKey"), expression.Name("Data"))
	//expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	expr, err := expression.NewBuilder().WithFilter(conditionOne.And(conditionTwo)).WithProjection(proj).Build()
	if err != nil {
		returnFormattedError(err)
	}

	queryParams := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TableName),
	}

	queryResult, err := dbSession().Scan(queryParams)
	if err != nil {
		returnFormattedError(err)
	}

	queryOut := make([]Pippo, len(queryResult.Items))
	for i, v := range queryResult.Items {
		tmp := Pippo{}
		dynamodbattribute.UnmarshalMap(v, &tmp)
		queryOut[i] = tmp
	}


	sess := session.Must(session.NewSession())
	svc := apigatewaymanagementapi.New(sess)

	log.Print(fmt.Sprintf("Active connections: %d", len(queryOut)))
	for _, v := range queryOut {
		connectionId := v.PrimaryKey[5:len(v.PrimaryKey)]
		connectionUrl := fmt.Sprintf("https://lquoki7y58.execute-api.ap-southeast-2.amazonaws.com/test/@connections/%s", connectionId)
		log.Print(fmt.Sprintf("Connection ID: %s", connectionId))
		log.Print(fmt.Sprintf("Connection URL: %s", connectionUrl))
		connectionString := aws.String(connectionUrl)
		params := apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: connectionString,
			Data: []byte("{\"hallo\": \"world\"}"),
		}
		output, err := svc.PostToConnection(&params)
		if err != nil {
			returnFormattedError(err)
		}
		log.Print(fmt.Sprintf("PostToConnection OUTPUT: %s", output.String()))
	}


	body, err := json.Marshal(queryOut)
	if err != nil {
		returnFormattedError(err)
	}
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 201}, nil
}

func apiGatewaySession() (apiGatewaySession apigatewayiface.APIGatewayAPI) {
	sess := session.Must(session.NewSession())

	return apigateway.New(sess)
}

func dbSession() (dbSession dynamodbiface.DynamoDBAPI) {
	sess := session.Must(session.NewSession())

	return dynamodb.New(sess)
}

func returnFormattedError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 422}, nil
}

func buildMessage(data string) (message Message) {
	return Message{
		PrimaryKey:		fmt.Sprintf("MESSAGE_%s", guuid.New().String()),
		SecondaryKey:	fmt.Sprintf("CHAT_%s", "42"),
		CreatedAt: 		time.Now().Format(time.RFC3339),
		Data:      		data,
	}
}
