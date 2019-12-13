package repos

import "github.com/claudioontheweb/lunch-app/pkg/db/models"

type UserRepository interface {
	GetAll() ([]*models.User, error)
	GetById(id string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) error
	CheckUser(email string) (bool, *models.User)
}