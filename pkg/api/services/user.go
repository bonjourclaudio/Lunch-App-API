package services

import (
	"github.com/claudioontheweb/lunch-app/pkg/api/repos"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	GetAllUsers() []*models.User
	GetUserById(id string) *models.User
	CreateUser(user *models.User) *models.User
	UpdateUser(user *models.User) *models.User
	DeleteUser(id string) error
	CheckUser(email string) (bool, *models.User)
}

type userService struct {
	repo repos.UserRepository
}

func NewUserService(r repos.UserRepository) UserService {
	return  &userService{
		repo: r,
	}
}

func (s *userService) GetAllUsers() []*models.User {
	users, err := s.repo.GetAll()
	if err != nil {
		log.Fatal("Failed to get Users from repo")
	}

	return users
}

func (s *userService) GetUserById(id string) *models.User {
	log.Debug("Getting User with id ", id)

	user, err := s.repo.GetById(id)
	if err != nil {
		log.Fatal("Failed getting User from repo")
	}

	log.Debug(user)

	return user
}

func (s *userService) CreateUser(user *models.User) *models.User {

	log.Debug("Creating new User with id ", user.ID)

	u, err := s.repo.Create(user)
	if err != nil {
		log.Fatal("Failed creating new user")
	}

	return u

}

func (s *userService) UpdateUser(user *models.User) *models.User {

	log.Debug("Updating User with ID ", user.ID)

	user, err := s.repo.Update(user)
	if err != nil {
		log.Fatal("Failed to update User")
	}

	return user

}

func (s *userService) DeleteUser(id string) error {

	log.Debug("Deleting User with ID ", id)

	err := s.repo.Delete(id)

	return err
}

func (s *userService) CheckUser(email string) (bool, *models.User) {

	log.Debug("Checking if User exists with Email ", email)
	exists, user := s.repo.CheckUser(email)

	return exists, user

}
