# go-playground-api
Go playground Web API using [Gin](https://github.com/gin-gonic/gin) and [Swaggo](https://github.com/swaggo/gin-swagger)

<p align="center">
  <img height="200" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
  <img height="200" src="https://avatars0.githubusercontent.com/u/29616670?s=400&v=4">
</p>

## Prerequisites
- [Install Go](https://golang.org/)
- [VS Code Extension for Go](https://github.com/microsoft/vscode-go)

## Installation
1. Download Gin framework: `go get github.com/gin-gonic/gin`
2. Download Swag for Go: `go get -u github.com/swaggo/swag/cmd/swag`
3. Download [gin-swagger](https://github.com/swaggo/gin-swagger) by using:
```sh
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```

## How to Run it
- From the terminal: `go run playground`
- Once it's running, go to a browser
- Hit the following url localhost:8080/swagger/index.html

## Generating Swagger documentation
If you change or add new apis:
1. Make sure they have the proper [General API](https://github.com/swaggo/swag/blob/master/README.md#general-api-info) annotations
2. Run the Swaggo command to regenerate the documentation: `swag init`
3. That's all :smiley:

## What you will find so far :heavy_check_mark:
- [x] Web API created with [Gin](https://github.com/gin-gonic/gin)
- [x] Swagger documentation of the APIs usign [Swaggo](https://github.com/swaggo/gin-swagger)
