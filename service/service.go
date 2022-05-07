package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/config"
)

func Init() {
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		config.GenerateConfigFile("config/config.example.yaml")
		panic(err)
	}
	if err := model.LoadDatabase(); err != nil {
		panic(err)
	}
}
