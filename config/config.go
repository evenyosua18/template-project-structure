package config

import (
	"fmt"
	"github.com/evenyosua18/template-project-structure/config/database"
	"github.com/evenyosua18/template-project-structure/config/server"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type configuration struct {
	Database database.ListDatabase
	Server   server.ConfigServer
}

const (
	serverLocation   = `\server\server.yml`
	databaseLocation = `\database\database.yml`
)

var (
	cfg configuration
)

func init() {
	log.Println("initialize configuration")

	//set variable
	var err error
	var location string

	//get location
	_, b, _, _ := runtime.Caller(0)
	location = filepath.Dir(b)

	//import server.yml
	var serverFile []byte
	var serverConf server.ConfigServer

	if serverFile, err = ioutil.ReadFile(location + serverLocation); err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(serverFile, &serverConf); err != nil {
		panic(err)
	}

	cfg.Server = serverConf

	//import database.yml
	var databaseFile []byte
	var databaseConf database.ListDatabase

	if databaseFile, err = ioutil.ReadFile(location + databaseLocation); err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(databaseFile, &databaseConf); err != nil {
		panic(err)
	}

	cfg.Database = databaseConf

	fmt.Println(cfg)
}

func GetConfig() *configuration {
	return &cfg
}
