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

	//request do
	user := &models.User{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		rw.Write([]byte("parse req body err"))
		log.Println(err)
		return
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	user.Id = 0
	user.Type = "user"

	//save
	err = store.User.Save(user)
	if err != nil {
		log.Println(err)
		rw.Write([]byte("db error"))
		return
	}

	//respose do
	output, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write([]byte(output))
}

func ShowUsersHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("show a user"))
}
