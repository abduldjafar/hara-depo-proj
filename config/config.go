package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	Postgres postgres
	Server   server
	Twillio  twillio
	Zensiva  zensiva
	Redis    redis
	Gcp      gcp
	Aws aws
}

type postgres struct {
	Url      string
	Ip       string
	Port     string
	Db       string
	User     string
	Password string
}

type server struct {
	Port string
}

type twillio struct {
	AuthToken  string
	AccountSid string
	From       string
}

type zensiva struct {
	Userkey string
	Passkey string
}

type redis struct {
	Host string
	Port string
}

type gcp struct {
	Bucket    string
	Projectid string
	Filename  string
}

type aws struct {
	Region string
	Bucketname string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		fmt.Println(err)
	}
}
