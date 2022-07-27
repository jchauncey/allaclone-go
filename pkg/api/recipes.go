package api

import (
	"allaclone/pkg/db"
	"allaclone/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sort"
)

func RecipeQueryHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling recipe request")
	vars := mux.Vars(r)

	db, err := db.Connect("root", "password", "localhost", "3306", "projecteq")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var recipes []models.Recipe
	if vars["name"] != "" {
		log.Debugf("query recipe by name - %s", vars["name"])
		db.Where("name LIKE ?", "%"+vars["name"]+"%").Find(&recipes)
	}

	if vars["tradeskill"] != "" && vars["min_skill"] == "" && vars["max_skill"] == "" {
		log.Debugf("querying %s recipes", vars["tradeskill"])
		db.Where("tradeskill = ?", models.GetTradeskillID(vars["tradeskill"])).Find(&recipes)
	}

	if vars["tradeskill"] != "" && vars["min_skill"] != "" && vars["max_skill"] == "" {
		log.Debugf("querying %s recipes with min (%s) skill values", vars["tradeskill"], vars["min_skill"])
		db.Where("tradeskill = ? AND trivial > ?", models.GetTradeskillID(vars["tradeskill"]), vars["min_skill"]).Find(&recipes)
	}

	if vars["tradeskill"] != "" && vars["min_skill"] != "" && vars["max_skill"] != "" {
		log.Debugf("querying %s recipes with min (%s) and max (%s) skill values", vars["tradeskill"], vars["min_skill"], vars["max_skill"])
		db.Where("tradeskill = ? AND trivial > ? AND trivial < ?", models.GetTradeskillID(vars["tradeskill"]), vars["min_skill"], vars["max_skill"]).Find(&recipes)
	}

	if vars["tradeskill"] != "" && vars["min_skill"] == "" && vars["max_skill"] != "" {
		log.Debugf("querying %s recipes with max (%s) skill values", vars["tradeskill"], vars["max_skill"])
		db.Where("tradeskill = ? AND trivial < ?", models.GetTradeskillID(vars["tradeskill"]), vars["max_skill"]).Find(&recipes)
	}

	sort.SliceStable(recipes, func(i, j int) bool {
		return recipes[i].Trivial < recipes[j].Trivial
	})

	data, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	_, _ = w.Write(data)
}
