package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/showntop/tantan/models"
)

func ListRelationshipsHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("list a user's all relationships")
	vars := mux.Vars(req)
	actorId, _ := strconv.Atoi(vars["user_id"])
	relationships, err := store.Relationship.FindAllByActorId(actorId)
	if err != nil {
		log.Panicln(err)
		rw.Write([]byte("db error"))
		return
	}

	output, err := json.Marshal(relationships)
	if err != nil {
		log.Panicln(err)
	}
	rw.Write([]byte(output))
}

func UpdateRelationshipsHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("update a user's relationship")
	vars := mux.Vars(req)
	actorId, _ := strconv.Atoi(vars["user_id"])
	relatorId, _ := strconv.Atoi(vars["other_user_id"])
	if actorId == relatorId {
		rw.Write([]byte("there is no need to do this"))
		return
	}
	//request do
	relationship := models.Relationship{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&relationship)
	if err != nil {
		rw.Write([]byte("parse req body err"))
		log.Println(err)
		return
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	relationship.Id = 0
	relationship.ActorId = actorId
	relationship.RelatorId = relatorId
	relationship.Type = "relationship"

	log.Println(relationship.State)
	err = relationship.Validate()
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}
	err = store.Relationship.Update(&relationship)
	if err != nil {
		log.Println(err)
		rw.Write([]byte("db error"))
		return
	}
	output, err := json.Marshal(relationship)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write([]byte(output))
}
