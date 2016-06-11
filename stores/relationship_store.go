package stores

import (
	"github.com/showntop/tantan/models"
)

type RelationshipStore struct {
	*Store
}

//find the record by actor and relator's id
func (r *RelationshipStore) FillByActorAndRelator(relatonship *models.Relationship) error {
	//relator's relationship
	err := r.Master.Model(relatonship).
		Where("actor_id = ? and relator_id = ?", relatonship.ActorId, relatonship.RelatorId).
		Select()
	if err.Error() == "pg: no rows in result set" {
		return nil
	}
	return err
}

func (r *RelationshipStore) Update(actorRelationship *models.Relationship) error {
	//
	relatorRelationship := actorRelationship.Reverse()
	err := r.FillByActorAndRelator(relatorRelationship)
	if err != nil {
		return err
	}
	//get the right state
	if relatorRelationship.State == "like" && actorRelationship.State == "like" {
		actorRelationship.State = "mathed"
		relatorRelationship.State = "mathed" //and update the relator
	}

	if relatorRelationship.State == "mathed" && actorRelationship.State == "dislike" {
		relatorRelationship.State = "like" //and update the relator
	}

	//update they then
	r.updateOne(actorRelationship)
	r.updateOne(relatorRelationship)
	return nil
}

//update the relationship
func (r *RelationshipStore) updateOne(relationship *models.Relationship) error {
	_, err := r.Master.Model(relationship).
		OnConflict("(actor_id, relator_id) DO UPDATE").
		Set("state = ?state").
		Create()
	return err
}
