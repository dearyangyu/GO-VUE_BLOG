package core

import (
	config "Go-blog-server/config/enter"
	"Go-blog-server/global"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatal("Config Init Unmarshal: %s", err)
	}
	log.Println("Config yamlFile load Init success.")
	// fmt.Println(c)
	global.Config = c
}
