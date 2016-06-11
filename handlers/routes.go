package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/showntop/tantan/stores"
)

var (
	store *stores.Store
)

func init() {

	store = stores.NewStore()
	fmt.Println(store)
	fmt.Println("storestorestorestorestorestorestorestorestorestore")

	r := mux.NewRouter()

	r.Methods("Get").Path("/users").HandlerFunc(ListUsersHandler)    //list all users
	r.Methods("Post").Path("/users").HandlerFunc(CreateUsersHandler) //create a user

	s := r.PathPrefix("/users").Subrouter()
	s.Methods("Get").Path("/{user_id}").HandlerFunc(ShowUsersHandler)                                         //show a user's info
	s.Methods("Get").Path("/{user_id}/relationships").HandlerFunc(ListRelationshipsHandler)                   //list a user's all relationships
	s.Methods("Put").Path("/{user_id}/relationships/{other_user_id}").HandlerFunc(UpdateRelationshipsHandler) //update the relationship of user

	log.Fatal(http.ListenAndServe(":8000", r))
}
