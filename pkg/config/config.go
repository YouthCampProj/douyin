package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Conf *Config
)

type Config struct {
	Debug bool `yaml:"debug"` // 是否开启调试模式
	MySQL struct {
		Host string `yaml:"host"` // Mysql地址
		Port int    `yaml:"port"` // Mysql端口
		User string `yaml:"user"` // Mysql用户名
		Pass string `yaml:"pass"` // Mysql密码
		Name string `yaml:"name"` // Mysql数据库名
	} `yaml:"mysql"`
	Salt string `yaml:"salt"` // 加密盐
	v    *viper.Viper
}

// LoadConfig 加载全局配置文件
func LoadConfig(path string) error {
	c := &Config{}
	c.v = viper.New()
	c.v.SetConfigFile(path)
	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	if err := c.v.Unmarshal(c); err != nil {
		return err
	}
	Conf = c
	return nil
}

// GenerateConfigFile 生成示例配置文件
func GenerateConfigFile(path string) {
	c := &Config{}
	c.Debug = true
	c.MySQL.Host = "localhost"
	c.MySQL.Port = 3306
	c.MySQL.User = "douyin"
	c.MySQL.Pass = "douyin"
	c.MySQL.Name = "douyin"
	c.Salt = "douyin"
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Println("生成示例配置文件失败：", err)
		return
	}
	if err := ioutil.WriteFile(path, data, 0600); err != nil {
		log.Println("生成示例配置文件失败：", err)
		return
	}
	log.Println("请按照" + path + "的格式填写配置文件")
}
