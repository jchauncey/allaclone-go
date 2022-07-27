package api

import (
	"allaclone/pkg/db"
	"allaclone/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func ListZonesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect("root", "password", "localhost", "3306", "projecteq")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var zones []models.Zone
	_ = db.Find(&zones)

	data, err := json.MarshalIndent(zones, "", "  ")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	_, _ = w.Write(data)
}

func ListZonesByPopulationHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect("root", "password", "localhost", "3306", "projecteq")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var zones []models.Zone
	_ = db.Find(&zones)

	for idx, z := range zones {
		var count int64
		db.Model(&models.Spawn{}).Where("zone = ?", z.ShortName).Count(&count)
		log.Debugf("Population count for %s - %d", z.ShortName, count)
		zones[idx].PopulationCount = count
	}

	data, err := json.MarshalIndent(zones, "", "  ")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	_, _ = w.Write(data)
}

func ZoneHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling zone request")
	vars := mux.Vars(r)

	db, err := db.Connect("root", "password", "localhost", "3306", "projecteq")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var zones []models.Zone
	if vars["id"] != "" {
		log.Debugf("query zone by id - %s", vars["id"])
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		db.Where(&models.Zone{ID: id}).Find(&zones)
	}

	if vars["shortname"] != "" {
		log.Debugf("query zone by short name - %s", vars["shortname"])
		db.Where(&models.Zone{ShortName: vars["shortname"]}).Find(&zones)
	}

	if vars["longname"] != "" {
		log.Debugf("query zone by long name - %s", vars["longname"])
		db.Where(&models.Zone{LongName: vars["longname"]}).Find(&zones)
	}

	if vars["expansion"] != "" {
		log.Debugf("query zone by expansion - %s", vars["expansion"])
		expansion, err := strconv.Atoi(vars["expansion"])
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		db.Where(&models.Zone{Expansion: expansion}).Find(&zones)
	}

	data, err := json.MarshalIndent(zones, "", "  ")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	_, _ = w.Write(data)
}
