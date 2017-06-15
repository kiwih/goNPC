package npcgen

//RaceTraits are for modifing NPCs based on their race
type RaceTraits struct {
	Speed int
	Size  string
	Type  string

	AbilityScoreIncreases AbilityScores
	AdultAge              int
	MaxAge                int

	RacialFeatures []Feature
}

var Races map[string]RaceTraits
