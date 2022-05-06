package service

import (
	"github.com/YouthCampProj/douyin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

// LoadDatabase 通过gorm加载数据库
func LoadDatabase() (*gorm.DB, error) {
	dsn := Conf.MySQL.User + ":" + Conf.MySQL.Pass + "@tcp(" + Conf.MySQL.Host + ":" + strconv.Itoa(Conf.MySQL.Port) + ")/" + Conf.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(
		&model.Comment{},
		&model.Favorite{},
		&model.Relation{},
		&model.Token{},
		&model.User{},
		&model.Video{},
	); err != nil {
		return nil, err
	}
	return db, nil
}
