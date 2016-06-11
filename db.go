package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
	"gopkg.in/pg.v4"
)

var queries = []string{
	`CREATE TABLE users
            (
              id serial NOT NULL,
              name character varying,
              CONSTRAINT users_pkey PRIMARY KEY (id)
            );`,

	`CREATE TYPE relationship_states AS ENUM ('like', 'dislike','matched');`,

	`CREATE TABLE relationships
            (
              id serial NOT NULL,
              actor_id serial NOT NULL,
              relator_id serial NOT NULL,
              state relationship_states,
              UNIQUE (actor_id, relator_id),
              CONSTRAINT relationships_pkey PRIMARY KEY (id),
              CONSTRAINT relationships_fkey FOREIGN KEY (actor_id)
                  REFERENCES users (id) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE CASCADE,
              CONSTRAINT relationships_fkey2 FOREIGN KEY (relator_id)
                  REFERENCES users (id) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE CASCADE
            );`,
}

func createSchema(db *pg.DB) error {
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "showntop",
		Password: "1",
		Database: "tantan2",
		// Whether to use secure TCP/IP connections (TLS).
		SSL: false,
	})
	pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))
	createSchema(db)
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
