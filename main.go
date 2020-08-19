package main

//go:generate go generate ./ent

import (
	"context"
	"log"

	"github.com/mattn/entgo-example/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	// run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	user := db.User.Create()
	user.SetName("mattn")
	user.SetAge(18)
	_, err = user.Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
