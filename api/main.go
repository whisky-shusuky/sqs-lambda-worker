package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type Body struct {
	Parameter string `json:"parameter"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// queryを取得
	parameter := aws.String(request.QueryStringParameters["parameter"])

	var resp events.APIGatewayProxyResponse
	body := Body{
		Parameter: *parameter,
	}
	jsonData, _ := json.Marshal(body)
	resp.StatusCode = 200
	resp.Body = string(jsonData)

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
