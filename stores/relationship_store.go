package stores

import (
	"github.com/showntop/tantan/models"
)

type RelationshipStore struct {
	*Store
}

//find a actor's relationship
func (r *RelationshipStore) FindAllByActorId(actorId int) ([]models.Relationship, error) {

	var relationships []models.Relationship
	err := r.Master.Model(&relationships).
		Where("actor_id = ?", actorId).
		Select()
	return relationships, err
}

//find the record by actor and relator's id
func (r *RelationshipStore) FillByActorAndRelator(relatonship *models.Relationship) error {
	//relator's relationship
	err := r.Master.Model(relatonship).
		Where("actor_id = ? and relator_id = ?", relatonship.ActorId, relatonship.RelatorId).
		Select()
	if err != nil && err.Error() == "pg: no rows in result set" {
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
		actorRelationship.State = "matched"
		relatorRelationship.State = "matched" //and update the relator
	}

	if relatorRelationship.State == "matched" && actorRelationship.State == "dislike" {
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
