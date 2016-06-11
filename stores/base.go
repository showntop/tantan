package stores

import (
	// "fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/pg.v4"
)

type Store struct {
	Master       *pg.DB   //*sqlx.DB
	Replicas     []*pg.DB //*sqlx.DB
	User         *UserStore
	Relationship *RelationshipStore
}

func NewStore(config map[string]string) *Store {
	store := &Store{}
	store.Master = setupDB(config)

	store.User = &UserStore{store}
	store.Relationship = &RelationshipStore{store}

	return store
}

func setupDB(config map[string]string) *pg.DB {

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
	return db
}
