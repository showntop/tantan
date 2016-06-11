package stores

import (
	// "fmt"
	"log"
	"os"

	// "github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
	"gopkg.in/pg.v4"
)

type Store struct {
	Master       *pg.DB   //*sqlx.DB
	Replicas     []*pg.DB //*sqlx.DB
	User         *UserStore
	Relationship *RelationshipStore
}

func NewStore() *Store {
	store := &Store{}
	store.Master = setupDB()

	store.User = &UserStore{store}
	store.Relationship = &RelationshipStore{store}

	return store
}

func setupDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "showntop",
		Password: "1",
		Database: "tantan2",
		// Whether to use secure TCP/IP connections (TLS).
		SSL: false,
	})
	pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))
	return db
}
