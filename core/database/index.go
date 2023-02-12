package database

import (
	"context"
	"time"

	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructDatabase)

var DB *mongo.Client

var MapTypeModel = map[string]string{
	constants.CustomerModel: "customers",
}

func Init(uri string) *mongo.Client {
	logger.Info("Init Database...With URI = ", "["+uri+"]")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbCon, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error("Connect Database Failed With Err= ", err.Error())
		panic(err)
	}
	DB = dbCon

	logger.Info("Connection Database Success...")
	return dbCon
}

func GetDB() *mongo.Database {
	return DB.Database("fast-food")
}

func GetCollection(schemaType string) *mongo.Collection {
	colName := MapTypeModel[schemaType]
	logger.Info("Collection Name = ", "["+colName+"]")
	return GetDB().Collection(colName)
}

// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.Find
func GetAll(schemaType string, filter interface{}, opts interface{}) (interface{}, error) {
	logger.Debug("[BEGIN] Datastore GetAll... With SchemaType = ", schemaType)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var findOpts *options.FindOptions

	opt := findOpts.SetSkip(0).SetLimit(10).SetSort(bson.D{{"createdAt", 1}})

	cursor, err := GetCollection(schemaType).Find(ctx, filter, opt)

	var results []bson.M

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	logger.Debug("[END] Func Create...")
	return results, err
}

func GetOne(schemaType string, filter interface{}, opts interface{}) {}

// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.InsertOne
var CreateOne = func(schemaType string, data interface{}) (*mongo.InsertOneResult, error) {
	logger.Debug("[BEGIN] Datastore CreateOne... With SchemaType = ", schemaType)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := GetCollection(schemaType).InsertOne(ctx, data)

	logger.Debug("[END] Datastore Create...")
	return res, err
}

func UpdateOne(schemaType string, filter interface{}, opts interface{}) {}

// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.CountDocuments
var Count = func(schemaType string, filter interface{}) int {
	logger.Debug("[BEGIN] Datastore Count... With SchemaType = ", schemaType)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, _ := GetCollection(schemaType).CountDocuments(ctx, filter)

	logger.Debug("[END] Datastore Count...")
	return int(count)
}
