package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host		string
	Port		string
	Dialect		string
	Debug		bool
	DBName		string
	DBUser		string
	DBPass		string
}

type ServerConfig struct {
	Host		string
	Port		string
	JWTSecret	string
}

type GoogleConfig struct {
	ClientID		string
	ClientSecret	string
	RedirectURL		string
	Scopes			[]string
}


type Config struct {
	DB     	DBConfig
	Server 	ServerConfig
	Google	GoogleConfig
}


func GetConfig() (*Config, error) {

	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading in Config file: " + err.Error())
	}

	dbConf := DBConfig{
		Host:  		viper.GetString("DB.HOST"),
		Port:   	viper.GetString("DB.PORT"),
		Dialect:	viper.GetString("DB.DIALECT"),
		Debug:		viper.GetBool("DB.DEBUG"),
		DBName: 	viper.GetString("DB.NAME"),
		DBUser: 	viper.GetString("DB.USER"),
		DBPass: 	viper.GetString("DB.PASS"),
	}

	serverConf := ServerConfig{
		Host: viper.GetString("SERVER.HOST"),
		Port: viper.GetString("SERVER.PORT"),
		JWTSecret: viper.GetString("JWT.SECRET"),
	}

	googleConf := GoogleConfig{
		ClientID:     viper.GetString("GOOGLE.CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE.CLIENT_SECRET"),
		RedirectURL:  viper.GetString("GOOGLE.REDIRECT_URL"),
		Scopes:       viper.GetStringSlice("GOOGLE.SCOPES"),
	}

	conf := Config{
		DB:     dbConf,
		Server: serverConf,
		Google: googleConf,
	}

	return &conf, nil
}