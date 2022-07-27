package models

import "strings"

func (Recipe) TableName() string {
	return "tradeskill_recipe"
}

type Recipe struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Tradeskill  int
	SkillNeeded int
	Trivial     int
	NoFail      int
	Materials   []Item `gorm:"many2many:tradeskill_recipe_entries;"`
}

type Item struct {
	ID   int
	Name string
}

type TradeskillRecipeEntries struct {
	RecipeID int
	ItemID   int
}

func GetTradeskillID(tradeskill string) int {
	switch {
	case strings.EqualFold(tradeskill, "alchemy"):
		return 59
	case strings.EqualFold(tradeskill, "Baking"):
		return 60
	case strings.EqualFold(tradeskill, "Blacksmithing"):
		return 63
	case strings.EqualFold(tradeskill, "Brewing"):
		return 65
	case strings.EqualFold(tradeskill, "Fishing"):
		return 55
	case strings.EqualFold(tradeskill, "Fletching"):
		return 64
	case strings.EqualFold(tradeskill, "Jewelery Making"):
		return 68
	case strings.EqualFold(tradeskill, "Poison Making"):
		return 56
	case strings.EqualFold(tradeskill, "Pottery Making"):
		return 69
	case strings.EqualFold(tradeskill, "Research"):
		return 58
	case strings.EqualFold(tradeskill, "Tailoring"):
		return 61
	case strings.EqualFold(tradeskill, "Tinkering"):
		return 57
	default:
		return 0
	}
}
