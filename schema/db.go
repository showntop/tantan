package schema

import (
	"log"
	"os"
	"strconv"

	// "github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
	"gopkg.in/pg.v4"
)

//need to split it to a single file
var queries = []string{

	`DO $$
             BEGIN
                IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'relationship_states') THEN
                    CREATE TYPE relationship_states AS ENUM ('liked', 'disliked','matched');
                END IF;
             END
             $$;`,

	`CREATE TABLE IF NOT EXISTS users
            (
              id serial NOT NULL,
              name character varying,
              type character varying DEFAULT 'user',
              CONSTRAINT users_pkey PRIMARY KEY (id)
            );`,

	`CREATE TABLE IF NOT EXISTS relationships
            (
              id serial NOT NULL,
              actor_id serial NOT NULL,
              relator_id serial NOT NULL,
              state relationship_states,
              type character varying DEFAULT 'relationship',
              UNIQUE (actor_id, relator_id),
              CONSTRAINT relationships_pkey PRIMARY KEY (id),
              CONSTRAINT relationships_fkey FOREIGN KEY (actor_id)
                  REFERENCES users (id) MATCH SIMPLE
                  ON UPDATE RESTRICT ON DELETE CASCADE,
              CONSTRAINT relationships_fkey2 FOREIGN KEY (relator_id)
                  REFERENCES users (id) MATCH SIMPLE
                  ON UPDATE RESTRICT ON DELETE CASCADE
            );`,
}

func createSchema(db *pg.DB) error {
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func InitDb(config map[string]string) {
	sslmode, err := strconv.ParseBool(config["sslmode"])
	if err == nil {
		log.Println("It's not ok for type sslmode")
	}

	db := pg.Connect(&pg.Options{
		Addr:     config["addr"],
		User:     config["user"],
		Password: config["password"],
		Database: config["dbname"],
		// Whether to use secure TCP/IP connections (TLS).
		SSL: sslmode,
	})
	pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))
	createSchema(db)
	err = db.Close()
	if err != nil {
		panic(err)
	}
}
