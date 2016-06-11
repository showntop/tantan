package stores

import (
	"fmt"

	"github.com/showntop/tantan/models"
)

type UserStore struct {
	*Store
}

func (u *UserStore) Save(user *models.User) error {

	err := u.Master.Create(user)

	if err != nil {
		fmt.Println("save :", err.Error())
		return err
	}
	return nil
}

func (u *UserStore) FindAll() (*[]models.User, error) {
	var users []models.User

	err := u.Master.Model(&users).Select()

	if err != nil {
		return &users, err
	}
	return &users, nil
}
