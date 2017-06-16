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
	Size   int
	Speed  int
}

//CombineStatBlocks combines two statblocks together
func CombineStatBlocks(a, b StatBlock) StatBlock {
	return StatBlock{} //TODO
}
