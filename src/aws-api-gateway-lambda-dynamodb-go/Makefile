PROJECT_SCOPE = sl-faas
PROJECT_NAME = aws-api-gateway-lambda-dynamodb-go

include .faas

test:
	@ go test ./...

build:
	@ GOOS=linux go build -o ./dist/handler/create ./src/create
	@ GOOS=linux go build -o ./dist/handler/resolve ./src/resolve
