package main

import (
	"github.com/claudioontheweb/lunch-app/config"
	"github.com/claudioontheweb/lunch-app/pkg/api/auth"
	"github.com/claudioontheweb/lunch-app/pkg/api/http"
	"github.com/claudioontheweb/lunch-app/pkg/api/repos"
	"github.com/claudioontheweb/lunch-app/pkg/api/services"
	"github.com/claudioontheweb/lunch-app/pkg/db"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		panic("Error retrieving Config: " + err.Error())
	}

	var userRepo repos.UserRepository

	userRepo, err = db.ConnectDB(&conf.DB)
	if err != nil {
		panic("Error connecting to DB: " + err.Error())
	}

	userService := services.NewUserService(userRepo)

	oa := auth.NewOAuth(&conf.Google)
	authService := auth.NewAuthService(oa)

	server := http.NewServer(userService, authService, &conf.Server)

	server.RegisterHandlers()

	log.Debug("Server starting up on port ", conf.Server.Port)
	err = server.StartHTTP()
	if err != nil {
		panic("Error Starting HTTP Server: " + err.Error())
	}

}
