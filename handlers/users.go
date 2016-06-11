package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/showntop/tantan/models"
)

func ListUsersHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("list all users")
	users, err := store.User.FindAll()
	if err != nil {
		log.Println(err)
		rw.Write([]byte("db error"))
		return
	}
	output, _ := json.Marshal(users)
	rw.Write([]byte(output))
}

func CreateUsersHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("create a user")

	name := req.FormValue("name")

	user := &models.User{Name: name, Type: "user"}
	err := store.User.Save(user)
	if err != nil {
		log.Println(err)
		rw.Write([]byte("db error"))
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write([]byte(output))
}

func ShowUsersHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("show a user"))
}
