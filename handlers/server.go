package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/showntop/tantan/stores"
)

var (
	store *stores.Store
)

func Setup(config map[string]interface{}) {
	//init the store
	store = stores.NewStore(config["database"].(map[string]string))
	log.Println("the store started...")

	//init the mux
	r := mux.NewRouter()
	//add routes
	r.Methods("Get").Path("/users").HandlerFunc(ListUsersHandler)    //list all users
	r.Methods("Post").Path("/users").HandlerFunc(CreateUsersHandler) //create a user

	s := r.PathPrefix("/users").Subrouter()
	s.Methods("Get").Path("/{user_id}").HandlerFunc(ShowUsersHandler)                                         //show a user's info
	s.Methods("Get").Path("/{user_id}/relationships").HandlerFunc(ListRelationshipsHandler)                   //list a user's all relationships
	s.Methods("Put").Path("/{user_id}/relationships/{other_user_id}").HandlerFunc(UpdateRelationshipsHandler) //update the relationship of user

	serverPort := config["port"].(string)
	log.Println("http listened on" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}
