package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/showntop/tantan/models"
)

func ListRelationshipsHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("list a user's all relationships")

	rw.Write([]byte("list a user's all relationships"))
}

func UpdateRelationshipsHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("update a user's relationship")
	vars := mux.Vars(req)
	actorId, _ := strconv.Atoi(vars["user_id"])
	relatorId, _ := strconv.Atoi(vars["other_user_id"])
	state := req.FormValue("state")
	relationship := models.Relationship{ActorId: actorId, RelatorId: relatorId, State: state}
	err := store.Relationship.Update(&relationship)
	if err != nil {
		fmt.Println(err)
		rw.Write([]byte("db error"))
		return
	}
	rw.Write([]byte("update a user's relationship"))
}
