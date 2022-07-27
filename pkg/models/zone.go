package models

var (
	supportExpansions = map[string]int{
		"Classic":                  1,
		"Ruins of Kunark":          2,
		"Scars of Velious":         3,
		"Shadows of Luclin":        4,
		"Planes of Power":          5,
		"Legacy of Ykesha":         6,
		"Lost Dungeons of Norrath": 7,
		"Gates of Discord":         8,
		"Omens of War":             9,
	}
)

func (Zone) TableName() string {
	return "zone"
}

type Zone struct {
	ID              int `gorm:"primaryKey"`
	ShortName       string
	LongName        string
	Expansion       int
	PopulationCount int64 `gorm:"-:all"`
}

func (Spawn) TableName() string {
	return "spawn2"
}

type Spawn struct {
	ID   int `gorm:"primaryKey"`
	Zone string
}
