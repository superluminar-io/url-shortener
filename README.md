# URL Shortener

[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](https://github.com/superluminar-io/url-shortener/blob/master/LICENSE.md)

> As always, there are countless ways to implement something. Based on the [Serverless URL Shortener in Go](https://github.com/superluminar-io/godays-workshop), this repository contains a variation of different implementations to reach the desired gaol.

## Scope

There was a time, where everybody used URL shorteners like `bit.ly`, `TinyURL`, or `ow.ly`. The feature set of an URL shortener is perfect as an example how to create, host, and operate a web service. Everybody has a clear understanding of the scope and boundaries of this application/service.

A basic URL shortener contains the following components:

* Storage for mapping between `short` and `full` URL
* API to create new `short` URL for a `full` URL
* API to access the `full` URL by providing the `short` URL

## Implementations

* [**`aws-api-gateway-lambda-dynamodb-go`**](src/aws-api-gateway-lambda-dynamodb-go) - AWS, API Gateway, DynamoDB, Lambda w/ Go

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
