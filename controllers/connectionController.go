package controllers

import (
	"github.com/anteater2/api/models"
	"github.com/gin-gonic/gin"
)

type connectJSON struct {
	Name string `json:"name" binding:"required"`
}

// HandleConnect handles the registration of a new client
func HandleConnect(c *gin.Context) {
	var json connectJSON
	if err := c.Bind(&json); err != nil {
		c.String(500, "Binding failed")
	}

	newClient := models.NewClient(json.Name, c.ClientIP())
	models.AddClient(newClient)

	c.String(200, "Welcome "+json.Name+"!")
}

// HandleDisconnect handles the disconnection of a client
func HandleDisconnect(c *gin.Context) {
	var json connectJSON
	if err := c.Bind(&json); err != nil {
		c.String(500, "Binding failed")
	}

	newClient := models.NewClient(json.Name, c.ClientIP())
	models.RemoveClient(newClient)

	c.String(200, "Goodbye "+json.Name+"!")
}

// HandlePoll returns the splice of clients
func HandlePoll(c *gin.Context) {
	c.String(200, models.GetClients().String())
}
