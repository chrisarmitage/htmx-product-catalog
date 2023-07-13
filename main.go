package main

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := connectToDb()
	s := NewServer("localhost:8080", db)
	s.Run()
}

func connectToDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to database")
	}

	db.AutoMigrate(
		&Product{},
	)

	// db.Create(Product{Id: ulid.Make(), Name: "Hand Weapon", Price: 7})
	// db.Create(Product{Id: ulid.Make(), Name: "Dagger", Price: 4})
	// db.Create(Product{Id: ulid.Make(), Name: "Spear", Price: 5})

	return db
}
