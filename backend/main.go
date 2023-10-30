package main

import (
	"fmt"
	"github.com/RenanLourenco/go-rest-api/database"
	"github.com/RenanLourenco/go-rest-api/routes"
)




func main() {
	fmt.Println("Initialize project with Go")
	database.Connection()
	routes.HandleRequest()
}