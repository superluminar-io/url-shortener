package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type handler struct {
	DynamoDBTableName string
	DynamoDBClient    dynamodbiface.DynamoDBAPI
}

func (h *handler) run(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s := request.PathParameters["id"]

	result, err := h.DynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(h.DynamoDBTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(s)},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Unable to access data"}, nil
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Unable to resolve URL mapping"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 302, Headers: map[string]string{"Location": *result.Item["url"].S}}, nil
}
