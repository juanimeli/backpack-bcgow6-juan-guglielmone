package main

import (
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/cmd/server/routes"
	"github.com/juanimeli/backpack-bcgow6-juan-guglielmone/db/implementaciondb/c1_tt/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")
}
