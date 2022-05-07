package model

import (
	"github.com/YouthCampProj/douyin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var (
	DB *gorm.DB
)

// LoadDatabase 通过gorm加载数据库
func LoadDatabase() error {
	dsn := config.Conf.MySQL.User + ":" + config.Conf.MySQL.Pass + "@tcp(" + config.Conf.MySQL.Host + ":" + strconv.Itoa(config.Conf.MySQL.Port) + ")/" + config.Conf.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// TODO: 开发完成后应该移除下述自动合并代码
	if err := db.AutoMigrate(
		&Comment{},
		&Favorite{},
		&Relation{},
		&Token{},
		&User{},
		&Video{},
	); err != nil {
		return err
	}
	DB = db
	return nil
}
