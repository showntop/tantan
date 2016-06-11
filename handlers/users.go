package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/showntop/tantan/models"
)

func ListUsersHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("list all users")
	users, err := store.User.FindAll()
	if err != nil {
		fmt.Println(err)
		rw.Write([]byte("db error"))
		return
	}
	a, _ := json.Marshal(users)
	rw.Write([]byte(a))
}

func CreateUsersHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("create a user")

	name := req.FormValue("name")

	user := &models.User{Name: name}
	fmt.Println(store)
	fmt.Println(store.User)
	err := store.User.Save(user)
	if err != nil {
		fmt.Println(err)
		rw.Write([]byte("db error"))
		return
	}

	rw.Write([]byte("create a user"))
}

func ShowUsersHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("show a user"))
}
