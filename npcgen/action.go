package npcgen

//An Action is something that an NPC/PC can do when it has the appropriate thing that grants this action
//e.g. a Sword would give the sword action
type Action struct {
	Damage     DiceFunction
	DamageType string

	Range ActionRange
}

//TODO: Actions need to depend on ability scores, e.g. when casting a spell we use spellcasting ability, etc, thrown weapons vs other things etc

//ActionRange is how far an action could go. Use 5ft for Melee
type ActionRange struct {
	Melee      bool
	Reach      bool
	ShortRange int
	LongRange  int
}

//if Melee and short/longRange, it's probably a melee/thrown weapon
//if not Melee and short/longRange, it's probably a ranged weapon
