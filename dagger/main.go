package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// get OS parent directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)

	source := client.Container().
		From("ubuntu:jammy").
		WithDirectory("/meta-helm-chart", client.Host().Directory(parent), dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"dagger/"},
		})

	// install helm
	runner := source.WithWorkdir("/meta-helm-chart").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "curl", "git"}).
		WithExec([]string{"bash", "-c", "curl -fsSL -o /tmp/get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3"}).
		WithExec([]string{"bash", "-c", "chmod 700 /tmp/get_helm.sh"}).
		WithExec([]string{"bash", "-c", "cd /tmp && DESIRED_VERSION=v3.12.3 ./get_helm.sh"})

	// run helm lint
	out, err := runner.WithExec([]string{"helm", "lint", "-f", "example/values.yaml", "."}).
		Stderr(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
