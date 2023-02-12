## Build Go with Gin Framework And Mongodb

- [Build Go with Gin Framework And Mongodb](#build-go-with-gin-framework-and-mongodb)
- [Golang For Beginner](#golang-for-beginner)
- [Go get and Go mod tidy](#go-get-and-go-mod-tidy)
- [Dev](#dev)
- [Test](#test)
- [Clone](#clone)
- [Environment variable](#environment-variable)
- [Structure](#structure)

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
  
  - github.com/onsi/ginkgo/v2 is a testing framework for Go designed to help you write expressive tests
  
  - github.com/onsi/gomega is the Ginkgo BDD-style testing framework's preferred matcher library.


- Parse json

```javascript
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

## Test

```sh
$ cd tests && go test
```
## Clone

```sh
$ git clone https://github.com/NguyenPhuocMinh/go-rest-api.git
```

## Environment variable

E.g:

- APP_PORT=8080
- APP_MONGO_URI=mongodb://127.0.0.1:27017

## Structure

```sh
.
├── commons                            => Defined method commons
├── configs                            => Defined environment variable
├── constants                          => Defined constants
├── core
│   ├── database                       => Init database
│   └── logger                         => Init logger
├── helpers                            => Defined func helper
├── main.go                            => Main
├── middleware                         => Defined middleware
├── resources                          => Defined success and error code
├── src
│   ├── controllers
│   │   └── v1                         => Defined controller v1
│   ├── dtos                           => Defined dto
│   ├── mappers                        => Convert entity to dto
│   ├── models                         => Defined model
│   ├── routers
│   │   └── v1                         => Defined router v1
│   └── services
│       └── v1                         => Defined service v1
├── tests                              => Testing
└── utils                              => Defined func utils
├── README.md
```
