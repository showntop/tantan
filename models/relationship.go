package models

const (
	Like    string = "like"
	Dislike string = "dislike"
	Mathed  string = "matched"
)

type Relationship struct {
	Id        int    `json:"myName,omitempty"`
	ActorId   int    `json:"myName,omitempty"`
	RelatorId int    `json:"id"`
	State     string `json:"state"`
}

//new the relator's relationship
func (r *Relationship) Reverse() *Relationship {
	return &Relationship{
		ActorId:   r.RelatorId,
		RelatorId: r.ActorId,
	}

}
