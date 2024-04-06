package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/gin-gonic/gin"
	"io"
	"kinexon/containerruntime/app/services"
	"net/http"
)

func Info(c *gin.Context) {

	docker := services.Docker
	info, err := docker.Info(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": info,
	})

}

func ListContainer(c *gin.Context) {

	imageName := c.Query("imageName")
	containerName := c.Query("containerName")

	var f filters.Args
	if imageName != "" && containerName != "" {
		f = filters.NewArgs(filters.KeyValuePair{Key: "ancestor", Value: imageName}, filters.KeyValuePair{Key: "name", Value: containerName})
	} else if containerName != "" {
		fmt.Println(containerName)
		f = filters.NewArgs(filters.KeyValuePair{Key: "name", Value: containerName})
	} else if imageName != "" {
		f = filters.NewArgs(filters.KeyValuePair{Key: "ancestor", Value: imageName})
	} else {
		f = filters.NewArgs()
	}

	fmt.Println(f)

	docker := services.Docker
	containers, err := docker.ContainerList(context.Background(), container.ListOptions{All: true, Filters: f})
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": containers,
	})
}

func RestartContainer(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)
	docker := services.Docker
	err := docker.ContainerRestart(context.Background(), id, container.StopOptions{})
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Container restarted successfully",
	})
}

func RemoveContainer(c *gin.Context) {

	id := c.Param("id")
	docker := services.Docker
	err := docker.ContainerRemove(context.Background(), id, container.RemoveOptions{})
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Container removed successfully",
	})
}

func StopContainer(c *gin.Context) {

	id := c.Param("id")
	docker := services.Docker
	err := docker.ContainerStop(context.Background(), id, container.StopOptions{})
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Container stopped successfully",
	})

}

func StartStreamContainerStats(c *gin.Context) {

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	containerID := c.Param("id")
	docker := services.Docker
	statsChan, err := services.GetContainerStats(context.Background(), docker, containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//defer close(statsChan) // Ensure channel is closed on exit

	c.Stream(func(w io.Writer) bool {
		select {
		case stats, ok := <-statsChan:
			if !ok {
				return false // Channel closed, stop streaming
			}
			// Marshal stats to JSON
			data, _err := json.Marshal(stats)
			if _err != nil {
				fmt.Println("Error marshalling stats to JSON:", err)
				return true // Continue streaming, handle error later
			}
			c.SSEvent("stats", string(data)) // Send stats as SSE event
			return true
		case <-c.Done():
			return false // Client disconnected, stop streaming
		}
	})
}
