package models

import (
	"errors"
	"strings"
)

const (
	Like    string = "liked"
	Dislike string = "disliked"
	Match   string = "matched"
)

type Relationship struct {
	Id        int    `json:"-"`
	ActorId   int    `json:"-"`
	RelatorId int    `json:"user_id"`
	State     string `json:"state"`
}

func (r *Relationship) Validate() error {
	if strings.Compare(Like, r.State) != 0 && strings.Compare(Dislike, r.State) != 0 {
		return errors.New("not allowed state")
	}
	return nil
}

//new the relator's relationship
func (r *Relationship) Reverse() *Relationship {
	return &Relationship{
		ActorId:   r.RelatorId,
		RelatorId: r.ActorId,
	}

}
