package playground20220405_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName          = "playground"
	migrationSource = "file://./resources/migrations/"
)

func runTest(t *testing.T) { // nolint: deadcode
	t.Helper()

	dbHost := getEnv("MONGO_27017_HOST", "localhost")
	dbPort := getEnv("MONGO_27017_PORT", "27017")
	dsn := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		assert.FailNowf(t, "failed to connect to mongo: %s", err.Error())
	}

	defer client.Disconnect(context.Background()) // nolint: errcheck,contextcheck

	db := client.Database(dbName)
	col := db.Collection("customer")

	result, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		assert.FailNowf(t, "failed to count documents: %s", err.Error())
	}

	assert.Equal(t, int64(1), result)
}

func getEnv(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}

	return val
}
