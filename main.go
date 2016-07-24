package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"

	"github.com/gotokatsuya/lboard/config"
	"github.com/gotokatsuya/lboard/library/database"
	"github.com/gotokatsuya/lboard/library/database/scheme"
	"github.com/gotokatsuya/lboard/route"
)

func init() {
	config.Load()
	db := database.NewGorm()
	scheme.Sync(db)
}

func main() {
	gracefulSrv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":9000",
			Handler: route.Init(),
		},
	}
	log.Infof("Serving %s with pid %d", gracefulSrv.Addr, os.Getpid())
	gracefulSrv.ListenAndServe()
}
