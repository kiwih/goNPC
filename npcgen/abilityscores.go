package npcgen

//AbilityScore is a handy type which allows us to easily compute modifier values
type AbilityScore int

//AbilityScores are used to hold the default six scores
type AbilityScores struct {
	Str AbilityScore
	Dex AbilityScore
	Con AbilityScore
	Int AbilityScore
	Wis AbilityScore
	Cha AbilityScore
}

//Modifier allows us to compute the modifier value on a AbilityScore
func (a AbilityScore) Modifier() int {
	return 0
}

//StatBlock is to be used on items, NPCs, races etc to enable calculation of NPC / PC final stats
type StatBlock struct {
	AbilityScores

	TempHP int
	Speed  int
}

//CombineStatBlocks combines two statblocks together
func CombineStatBlocks(a, b StatBlock) StatBlock {
	x := StatBlock{}
	x.Str = a.Str + b.Str
	x.Dex = a.Dex + b.Dex
	x.Con = a.Con + b.Con
	x.Int = a.Int + b.Int
	x.Wis = a.Wis + b.Wis
	x.Cha = a.Cha + b.Cha

	if a.TempHP > b.TempHP {
		x.TempHP = a.TempHP
	} else {
		x.TempHP = b.TempHP
	}

	x.Speed = a.Speed + b.Speed

	return x
}
