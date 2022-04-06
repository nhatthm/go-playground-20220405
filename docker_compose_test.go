//go:build test_with_docker_compose

package playground20220405_test

import "testing"

func TestWithDockerCompose(t *testing.T) {
	t.Parallel()

	// =========================================================================
	// Run tests
	// =========================================================================
	runTest(t)
}
