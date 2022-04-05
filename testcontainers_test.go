//go:build test_with_testcontainers

package playground20220405_test

import (
	"context"
	"testing"
	"time"

	"github.com/nhatthm/testcontainers-go-registry/database/mongo"
	"github.com/stretchr/testify/assert"
)

func TestWithTestContainers(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("starting mongo container")

	_, err := mongo.StartGenericContainer(ctx,
		mongo.RunMigrations(migrationSource, dbName),
	)
	if err != nil {
		assert.FailNow(t, "failed to start mongo container: %s", err.Error())
	}

	runTest(t)
}
