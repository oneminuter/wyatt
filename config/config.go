package config

import (
	"log"
	"os"

	"f.in/v/utils"
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server server
	MGO    mongoInfo
	Redis  redisInfo
	Mysql  mysqlInfo
	Debug  bool
}

type server struct {
	Port int
}

type mysqlInfo struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DbName  string
	Charset string
	Pool    int
}

type redisInfo struct {
	Host   string
	Port   int
	Pass   string
	Pool   int
	DbName int
}

type mongoInfo struct {
	Host string
	Port int
	User string
	Pass string
	Pool int
}

func GetConfig() TomlConfig {
	debug := utils.BoolMust(os.Getenv("DEBUG"))
	tomlFile := utils.On(!debug, "config", "config_debug").(string)

	var config TomlConfig
	config.Debug = debug
	if _, err := toml.DecodeFile(utils.SelfDir()+"/config/"+tomlFile+".toml", &config); err != nil {
		log.Println(err)
		return config
	}
	return config
}
