package questgen

type BasicNPC struct {
	Name   string
	Gender string
	Race   string
}

func (b BasicNPC) GetName() string {
	return b.Name
}

func (b BasicNPC) GetRace() string {
	return b.Race
}

func (b BasicNPC) GetGender() string {
	return b.Gender
}

func GetBasicQuestNPCs() []QuestNPC {
	return []QuestNPC{
		BasicNPC{
			"Alfonzo McKerny",
			"Male",
			"Human",
		},
		BasicNPC{
			"Alys Wen",
			"Female",
			"Elf",
		},
	}
}

type BasicLocation string

func (b BasicLocation) GetName() string {
	return string(b)
}

func GetBasicQuestLocations() []QuestLocation {
	return []QuestLocation{
		BasicLocation(
			"The Folicking Foal Freehouse",
		),
		BasicLocation(
			"The Learned Library",
		),
		BasicLocation(
			"The Dark Grotto",
		),
	}
}

type BasicItem string

func (b BasicItem) GetName() string {
	return string(b)
}

func GetBasicQuestItems() []QuestItem {
	return []QuestItem{
		BasicItem(
			"Old Pearl Necklace",
		),
		BasicItem(
			"Dusty Book",
		),
		BasicItem(
			"Red-spotted Toadstool",
		),
	}
}
