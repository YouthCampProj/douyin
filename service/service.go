package service

import (
	"github.com/YouthCampProj/douyin/model"
	"github.com/YouthCampProj/douyin/pkg/config"
	"github.com/YouthCampProj/douyin/pkg/snowflake"
)

func Init() {
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		config.GenerateConfigFile("config/config.example.yaml")
		panic(err)
	}
	if err := model.LoadDatabase(); err != nil {
		panic(err)
	}
	// 雪花算法初始化
	if err := snowflake.Init("2022-06-06", 1); err != nil {
		panic(err)
	}
}
