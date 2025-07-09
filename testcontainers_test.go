//go:build test_with_testcontainers

package playground20220405_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.nhat.io/testcontainers-extra"
	"go.nhat.io/testcontainers-registry/mongo"
)

func TestWithTestContainers(t *testing.T) {
	t.Parallel()

	mongoVersion := getEnv("MONGO_VERSION", "6.0")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Logf("starting mongo:%s container", mongoVersion)

	_, err := mongo.StartGenericContainer(ctx,
		mongo.RunMigrations(migrationSource, dbName),
		testcontainers.WithImageTag(mongoVersion),
	)
	if err != nil {
		assert.FailNow(t, "failed to start mongo container: %s", err.Error())
	}

	// =========================================================================
	// Run tests
	// =========================================================================
	runTest(t)
}
