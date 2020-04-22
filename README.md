# go-playground-api
Go playground Web API using [Gin](https://github.com/gin-gonic/gin) and [Swaggo](https://github.com/swaggo/gin-swagger) (Go version 1.14)

<p align="center">
  <img height="200" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
  <img height="200" src="https://avatars0.githubusercontent.com/u/29616670?s=400&v=4">
</p>

## Notes
The idea of this project is to provide all the apis that exists in my other repository [NetCore.Playground.Api](https://github.com/joacod/NetCore.Playground.Api)

What we get from this:
- We practice different approaches in Go
- We can compare **Go** implementations with the same ones written in **.NetCore**
- We have two versions of the same Web Api written in different languages, that can be switched from a UI having the exact same results

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
1. From the terminal: `go run playground`
2. Once it's running, go to a browser
3. Hit the following url http://localhost:8080/swagger/index.html

## Generating Swagger documentation
If you change or add new apis:
1. Make sure they have the proper [General API](https://github.com/swaggo/swag/blob/master/README.md#general-api-info) annotations
2. Run the Swaggo command to regenerate the documentation: `swag init`
3. That's all :smiley:

## What you will find so far :heavy_check_mark:
- [x] Web API created with [Gin](https://github.com/gin-gonic/gin)
- [x] Swagger documentation of the APIs usign [Swaggo](https://github.com/swaggo/gin-swagger)
