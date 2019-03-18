package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSession())

	h := &handler{
		DynamoDBTableName: os.Getenv("DYNAMODB_TABLE_NAME"),
		DynamoDBClient:    dynamodb.New(sess),
	}

	lambda.Start(h.run)
}
