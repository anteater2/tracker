package main

import (
	"fmt"
	"os"

	"github.com/anteater2/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.HandlePoll)
	r.POST("/", controllers.HandleConnect)
	r.DELETE("/", controllers.HandleDisconnect)

	port, found := os.LookupEnv("PORT")
	if !found {
		port = ":5000"
	} else {
		port = ":" + port
	}
	fmt.Printf("Running on %s\n", port)

	r.Run(port)
}
