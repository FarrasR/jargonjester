package main

import (
	"jargonjester/database"
	"jargonjester/database/migration"
	"jargonjester/utils"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	utils.LoadEnvironment()

	ins := database.InitDB()

	m := gormigrate.New(ins, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migration.V20230410151447,
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
