package config

import (
	"github.com/go-ini/ini"
	"log"
	"novel/global"
)

type Database struct {
	Host        string `ini:"host"`
	Port        string `ini:"port"`
	Type        string `ini:"type"`
	User        string `ini:"user"`
	Password    string `ini:"password"`
	Database    string `ini:"database"`
	Charset     string `ini:"charset"`
	TablePrefix string `ini:"tablePrefix"`
}

var DbConfig = &Database{}

func Setup() {
	iniFile, err := ini.Load(global.RootPath + "/conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	err = iniFile.Section("database").MapTo(DbConfig)
	if err != nil {
		log.Fatalf("section database fail to parse: %v", err)
	}
}
