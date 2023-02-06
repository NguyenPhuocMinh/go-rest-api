## Build Go with Gin Framework And Mongodb

- [Build Go with Gin Framework And Mongodb](#build-go-with-gin-framework-and-mongodb)
- [Golang For Beginner](#golang-for-beginner)
- [Go get and Go mod tidy](#go-get-and-go-mod-tidy)
- [Dev](#dev)
- [Clone](#clone)
- [Environment variable](#environment-variable)
- [Ex .env](#ex-env)

## Golang For Beginner

- We need to initialize a Go module to manage project dependencies by running the command below:

```sh
$ go mod init fast-food-api-client
```

    - This command will create a go.mod file for tracking project dependencies.

- We proceed to install the required dependencies with:

```sh
$ go get -u github.com/gin-gonic/gin go.mongodb.org/mongo-driver/mongo github.com/joho/godotenv github.com/go-playground/validator/v10
```

- Dependencies info

  - github.com/gin-gonic/gin is a framework for building web applications.

  - go.mongodb.org/mongo-driver/mongo is a driver for connecting to MongoDB.

  - github.com/joho/godotenv is a library for managing environment variables.

  - github.com/go-playground/validator/v10 is a library for validating structs and fields.

- Format json

```go
  data, _ := json.MarshalIndent(res, "", " ")
	fmt.Println("hehe", string(data))
```

## Go get and Go mod tidy

- go get -u upgrade libraries to new versions
- go get -u module_name upgrade a specific library
- go mod tidy add required libraries, remove unused libraries, or upgrade libraries from older versions to the latest versions. go mod tidy is more versatile, easier to use than go get. It helps programmers have less to worry about.

## Dev

- Start

```sh
$ go run main.go
```

## Clone

```sh
$ git clone https://github.com/NguyenPhuocMinh/pm-tool-api.git
```

## Environment variable

## Ex .env

- APP_PORT=8080
- APP_DOCS_PATH=/docs
- APP_DOMAIN_PATH=http://localhost:3000
- APP_SECRET_KEY=do-biet-day-secret
- APP_AUDIENCE=pm-tool-aud
- APP_ISSUER=http://localhost:3000