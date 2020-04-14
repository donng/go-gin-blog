package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Database struct {
	Type        string `yaml:"Type"`
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	DBName      string `yaml:"DBName"`
	TablePrefix string `yaml:"TablePrefix"`
}

type Server struct {
	HttpPort     int `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

type Config struct {
	Database Database `yaml:"Database"`
	Server   Server   `yaml:"Server"`
}

var Setting = &Config{}

func init() {
	yamlFile, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to read yaml file, err: %s", err))
	}
	err = yaml.Unmarshal(yamlFile, &Setting)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve yaml file, err: %s", err))
	}

	Setting.Server.WriteTimeout = Setting.Server.WriteTimeout * time.Second
	Setting.Server.ReadTimeout  = Setting.Server.ReadTimeout * time.Second
}
