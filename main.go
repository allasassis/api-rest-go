package main

import (
	"github.com/allasassis/api-rest-go.git/models"
	"github.com/allasassis/api-rest-go.git/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Test", Story: "TestStory"},
		{Id: 2, Name: "Test2", Story: "TestStory2"},
	}
	routes.HandleRequest()
}
