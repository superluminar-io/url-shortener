# URL Shortener

> Part of [url-shortener](https://github.com/superluminar-io/url-shortener).

Using Amazon API Gateway, AWS Lambda, Amazon DynamoDB, and Go source code this folder showcases a serverless URL shortening service on AWS using the [AWS Serverless Application Model](https://github.com/awslabs/serverless-application-model).

## Configure

```bash
# Configure S3 Bucket to store artifacts
$ > make configure
```

### Deployment

```bash
# Build Go code, create CloudFormation Stack, and deploy it …
$ > make build package deploy
```

### Outputs

```bash
# Retrieve deployed API Gateway URL
$ > make outputs

[
  {
    "OutputKey": "URL",
    "OutputValue": "https://qgh8u9yro0.execute-api.eu-west-1.amazonaws.com/Prod",
    "Description": "URL for HTTPS Endpoint"
  }
]
```

## Usage

### Create URL

```bash
# Create new short URL for https://superluminar.io
$ > curl -X POST \
    https://qgh8u9yro0.execute-api.eu-west-1.amazonaws.com/Prod \
    --data '{"url": "https://superluminar.io"}'

1zzawj6g59ghf
```

### Resolve URL

```bash
# Access destination using short URL
$ > curl -i \
    https://qgh8u9yro0.execute-api.eu-west-1.amazonaws.com/Prod/1zzawj6g59ghf

HTTP/2 302
content-type: application/json
content-length: 0
location: https://superluminar.io
```