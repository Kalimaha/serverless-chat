package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	ResponseId string
}

// Invoke it locally
// aws --profile sideprojects lambda invoke --function-name RetrievePotatoes --payload '{"pathParameters": {"requestId": "42"}}' response.json
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestId	:= request.PathParameters["requestId"]
	response 	:= MyResponse{ResponseId: requestId}
	body, _ 	:= json.Marshal(response)

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
