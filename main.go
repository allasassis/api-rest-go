package main

import (
	"fmt"
	"github.com/allasassis/api-rest-go.git/database"
	"github.com/allasassis/api-rest-go.git/routes"
)

func main() {
	fmt.Println("Starting API Go...")
	database.ConnectWithDatabase()
	routes.HandleRequest()
}
