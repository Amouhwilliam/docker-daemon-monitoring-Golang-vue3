package services

import (
	"context"
	"github.com/docker/docker/client"
	"log/slog"
	"os"
)

var Docker *client.Client

func CreateDockerClient() {
	var err error
	Docker, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		slog.Error("failed to instantiate docker client: %v", err)
	}
	_, err = Docker.Ping(context.Background())
	if err != nil {
		slog.Error("failed to connect to docker daemon: %v", err)
		os.Exit(1)
	}
}
