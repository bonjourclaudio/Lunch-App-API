package mysql

import (
	"github.com/claudioontheweb/lunch-app/pkg/api/repos"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) repos.UserRepository {
	return &userRepository{
		db: db,
	}
}


func (r *userRepository) GetAll() ([]*models.User, error) {
	var res []*models.User
	if err := r.db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *userRepository) GetById(id string) (*models.User, error) {
	var res models.User
	if err := r.db.Find(&res, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	return user, r.db.Update(&user).Error
}

func (r *userRepository) Delete(id string) error {
	r.db.Where("id = ?", id).Delete(&models.User{})
	return r.db.Error
}


func (r *userRepository) CheckUser(email string) (bool, *models.User) {

	user := models.User{}

	if r.db.First(&user,"email = ?", email).RecordNotFound() {
		return false, nil
	} else {
		return true, &user
	}

}