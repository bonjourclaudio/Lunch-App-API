package db

import (
	"github.com/claudioontheweb/lunch-app/config"
	"github.com/claudioontheweb/lunch-app/pkg/api/repos"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	"github.com/claudioontheweb/lunch-app/pkg/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func ConnectDB(config *config.DBConfig) (repos.UserRepository, error) {

	switch config.Dialect {

	case "mysql":
		return connectMysql(config)

	default:
		log.Fatal("No valid DB dialect specified")
	}

	return nil, nil
}

func connectMysql(config *config.DBConfig) (repos.UserRepository, error) {

	db, err := gorm.Open(config.Dialect, config.DBUser + ":" +config.DBPass + "@/" + config.DBName +"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	db.LogMode(config.Debug)

	userRepo := mysql.NewMysqlUserRepository(db)

	return userRepo, nil

}

func autoMigrate(db *gorm.DB) error {
	log.Debug("Auto migrating DB Models")

	return db.AutoMigrate(&models.User{}).Error
}