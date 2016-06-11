package stores

import (
	"github.com/showntop/tantan/models"
)

type RelationshipStore struct {
	*Store
}

//find a actor's relationship
func (r *RelationshipStore) FindAllByActorId(actorId int) ([]models.Relationship, error) {

	relationships := []models.Relationship{}
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

//update or create a user's relation to others
func (r *RelationshipStore) Update(actorRelationship *models.Relationship) error {
	//get the relator's ship
	relatorRelationship := actorRelationship.Reverse()
	err := r.FillByActorAndRelator(relatorRelationship)
	if err != nil {
		return err
	}
	//get the right state
	if (relatorRelationship.State == models.Like || relatorRelationship.State == models.Match) && actorRelationship.State == models.Like {
		actorRelationship.State = models.Match
		relatorRelationship.State = models.Match //and update the relator
	}
	if relatorRelationship.State == models.Match && actorRelationship.State == models.Dislike {
		relatorRelationship.State = models.Like //and update the relator
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
