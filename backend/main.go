package main

import (
	"kinexon/containerruntime/app/services"
	"kinexon/containerruntime/utils"
)

func init() {
	//utils.LoadEnvVariables()
	services.CreateDockerClient()
}

func main() {
	utils.RunServer("8080")
}
