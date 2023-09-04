package main

import (
	"assignment02/databases"
	"assignment02/routes"
)

func main() {
	databases.StartDB()
	var PORT = ":8080"

	routes.StartServer().Run(PORT)

}
