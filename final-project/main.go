package main

import (
	"final-project/databases"
	"final-project/routes"
)

func main() {
	databases.StartDB()
	var PORT = ":8080"

	routes.StartServer().Run(PORT)

}
