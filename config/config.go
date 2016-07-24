package config

import (
	"path"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	libpath "github.com/gotokatsuya/lboard/library/path"
)

var c config

type config struct {
	DB   Database `toml:"database"`
	Time Time     `toml:"time"`
}

// Database ...
type Database struct {
	Driver string `toml:"driver"`
	Source string `toml:"source"`
}

// Time ...
type Time struct {
	Zone string `toml:"zone"`
}

// GetDB return database configuration
func GetDB() Database {
	return c.DB
}

// GetTime return time configuration
func GetTime() Time {
	return c.Time
}

// Load ...
func Load() {
	tmlPath := path.Join(libpath.GetAppPath(), "config", gin.Mode(), "config.tml")
	_, err := toml.DecodeFile(tmlPath, &c)
	if err != nil {
		log.WithFields(log.Fields{
			"Path":  tmlPath,
			"Error": err,
		}).Panicln("Config:Load toml.DecodeFile")
	}
}
