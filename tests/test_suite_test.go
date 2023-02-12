package tests_test

import (
	"context"
	"os"
	"testing"

	"fast-food-api-client/core/database"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

// see the document https://onsi.github.io/ginkgo/#getting-started

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = BeforeSuite(func() {
	var err error

	err = os.Setenv("APP_MONGO_URI", "mongodb://127.0.0.1:27017")

	// connect database
	DB = database.Init(os.Getenv("APP_MONGO_URI"))
	err = DB.Ping(context.Background(), readpref.Primary())
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	DB.Disconnect(context.Background())
})
