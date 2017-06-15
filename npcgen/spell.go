package npcgen

//SpellLevel is used to indicate a spell's level
type SpellLevel int

//Spell is used to store a D&D SRD spell
type Spell struct {
	Name                string
	Level               SpellLevel
	School              string
	IsRitual            bool
	CastingTime         string
	Range               string
	HasVerbal           bool
	HasSomatic          bool
	HasMaterial         bool
	MaterialComponent   string
	HasMaterialConsumed bool
	MaterialGPCost      string
	Duration            string
	IsConcentration     bool
	Description         string
	AtHigherLevels      string
}
