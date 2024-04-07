package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"kinexon/containerruntime/app/services"
	"net/http"
)

func ContainersList(c *gin.Context) {
	imageName := c.Query("imageName")
	containerName := c.Query("containerName")

	containers, err := services.ContainersList(&imageName, &containerName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": *containers,
	})
}

func RestartContainer(c *gin.Context) {
	id := c.Param("id")
	err := services.RestartContainer(&id)

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
	err := services.RemoveContainer(&id)

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
	err := services.StopContainer(&id)

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
	statsChan, err := services.GetContainerStats(context.Background(), containerID)
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
