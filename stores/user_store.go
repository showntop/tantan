package stores

import (
	"github.com/showntop/tantan/models"
)

type UserStore struct {
	*Store
}

func (u *UserStore) Save(user *models.User) error {

	err := u.Master.Create(user)
	return err
}

func (u *UserStore) FindAll() ([]models.User, error) {

	users := []models.User{}
	err := u.Master.Model(&users).Select()
	return users, err
}
