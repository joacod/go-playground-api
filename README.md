# Go Playground
Go playground Web API using [Gin](https://github.com/gin-gonic/gin), [Swaggo](https://github.com/swaggo/gin-swagger), [CORS gin's middleware](https://github.com/gin-contrib/cors) and [GORM](https://gorm.io/)

<p align="center">
  <img height="200" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
  <img height="200" src="https://avatars0.githubusercontent.com/u/29616670?s=400&v=4">
</p>

## Notes
Go version 1.14

The idea of this project is to provide all the apis that exists in my other repositories [NetCore.Playground.Api](https://github.com/joacod/NetCore.Playground.Api) and [node-playground-api](https://github.com/joacod/node-playground-api)

What we get from this:
- We practice different approaches in Go
- We can compare **Go** implementations with the same ones written in **.NetCore** and **Node.js**
- We have three versions of the same Web Api written in different languages, that can be switched from a UI having the exact same results

I have created three UI projects to consume the APIs created here, you can use either of the following repositories:
- [Angular Playground](https://github.com/joacod/angular-playground-ui)
- [Vue Playground](https://github.com/joacod/vue-playground-ui)
- [React Playground](https://github.com/joacod/react-playground-ui)

## Prerequisites
- [Install Go](https://golang.org/)
- [VS Code Extension for Go](https://github.com/microsoft/vscode-go)
- [MySQL instance in your system](https://www.mysql.com/)
    - This is optional and used only for implemented CRUD operations using GORM
    - If you are not interested in this part, you can avoid the installtion of MySQL and just change the flag named *useMysqlDb* in main.go file

## Installation and Setup
1. Download Gin framework: `go get github.com/gin-gonic/gin`
2. Download Swag for Go: `go get -u github.com/swaggo/swag/cmd/swag`
3. Download [gin-swagger](https://github.com/swaggo/gin-swagger) by using:
```sh
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```
4. Download MySQL driver for Go: `go get github.com/go-sql-driver/mysql`
5. Download GORM: `go get github.com/jinzhu/gorm`
6. Go to the file *database/Database.go* and update DBName, Host, Port, User and Password according to your local environment configuration

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
- [x] CRUD operations using GORM and MySQL
