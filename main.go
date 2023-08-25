package main

import (
	"context"
	"fmt"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types"
)

func main() {
	cmd, err := command.NewDockerCli()
	if err != nil {
		panic(err)
	}
	if err := cmd.Initialize(flags.NewClientOptions()); err != nil {
		panic(err)
	}
	c := cmd.Client()

	images, err := c.ImageList(context.TODO(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, i := range images {
		fmt.Println(i.ID)
	}
}
