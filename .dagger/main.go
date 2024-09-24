package main

import (
	"context"

	"dagger/health/internal/dagger"
)

type Health struct{}

// Build and publish image from existing Dockerfile
func (m *Health) Build(
	ctx context.Context,
	// location of directory containing Dockerfile
	src *dagger.Directory,
	// Registry address
	//+optional
	//+default: "awear/healthy"
	registry string,
	// Registry username
	//+optional
	//+default: "awear"
	username string,
	// Registry password
	password *dagger.Secret,
) (string, error) {
	ref, err := dag.Container().
		WithDirectory("/src", src).
		WithWorkdir("/src").
		Directory("/src").
		DockerBuild(). // build from Dockerfile
		WithRegistryAuth(registry, username, password).
		// Publish(ctx, "ttl.sh/health:1h")
		Publish(ctx, "docker.io/awear/health:latest")

	if err != nil {
		return "", err
	}

	return ref, nil
}
