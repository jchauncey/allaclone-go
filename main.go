package main

import (
	"allaclone/pkg/api"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)
	r := mux.NewRouter()

	r.HandleFunc("/recipes", api.RecipeQueryHandler).Queries("tradeskill", "{tradeskill}", "min_skill", "{min_skill}", "max_skill", "{max_skill}")
	r.HandleFunc("/recipes", api.RecipeQueryHandler).Queries("tradeskill", "{tradeskill}", "max_skill", "{max_skill}")
	r.HandleFunc("/recipes", api.RecipeQueryHandler).Queries("tradeskill", "{tradeskill}", "min_skill", "{min_skill}")
	r.HandleFunc("/recipes", api.RecipeQueryHandler).Queries("tradeskill", "{tradeskill}")
	r.HandleFunc("/recipes", api.RecipeQueryHandler).Queries("name", "{name}")

	r.HandleFunc("/zones", api.ListZonesHandler)
	r.HandleFunc("/zonesByPopulation", api.ListZonesByPopulationHandler)
	r.HandleFunc("/zone", api.ZoneHandler).Queries("id", "{id}")
	r.HandleFunc("/zone", api.ZoneHandler).Queries("shortname", "{shortname}")
	r.HandleFunc("/zone", api.ZoneHandler).Queries("expansion", "{expansion}")
	log.Info("Serving content on port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
