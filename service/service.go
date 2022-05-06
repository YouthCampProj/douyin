package service

import (
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Conf *Config
)

func Init() {
	var err error
	Conf, err = LoadConfig("config/config.yaml")
	if err != nil {
		GenerateConfigFile("config/config.example.yaml")
		panic(err)
	}
	DB, err = LoadDatabase()
	if err != nil {
		panic(err)
	}
}
