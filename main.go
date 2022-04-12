package main

import (
	"allaclone/pkg/db"
	"allaclone/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {

	db, err := db.Connect("root", "password", "localhost", "3306", "projecteq")
	if err != nil {
		log.Fatal(err)
	}

	var zones []models.Zone
	_ = db.Find(&zones)

	for _, z := range zones {
		log.Infof("Zone: %+v", z)
	}

}
