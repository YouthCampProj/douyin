package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

// LoadDatabase 通过gorm加载数据库
func LoadDatabase() (*gorm.DB, error) {
	dsn := Conf.MySQL.User + ":" + Conf.MySQL.Pass + "@tcp(" + Conf.MySQL.Host + ":" + strconv.Itoa(Conf.MySQL.Port) + ")/" + Conf.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
