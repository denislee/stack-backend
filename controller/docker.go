package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"github.com/denislee/stack-backend/config"
)

type DockerController struct{}

func (p DockerController) Status(c *gin.Context) {

	dockerApiVersion := config.GetDockerApiVersion() 
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerApiVersion))
	if err != nil {
		panic(err)
	}

	status, err := cli.Status()
	if err != nil {
		panic(err)
	}

	c.JSON(200, status)

	return
}

func (p DockerController) Containers(c *gin.Context) {

	dockerApiVersion := config.GetDockerApiVersion() 
	cli, err := client.NewClientWithOpts(client.WithVersion(dockerApiVersion))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}

	c.JSON(200, containers)

	return
}
