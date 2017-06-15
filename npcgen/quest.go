package npcgen

//QuestType provides the nomenclature for the type of Quest trope
type QuestType int

const (
	//QuestTypeKill is for being sent to kill something, e.g. Barkeep wants you to kill the rats in his cellar
	QuestTypeKill QuestType = iota

	//QuestTypeDelivery is for being sent to deliver an item, e.g. Barkeep needs awkwardly large beer delivery sent some distance
	QuestTypeDelivery

	//QuestTypeCollection is for being sent to collect an item and return it, e.g. Barkeep needs hops from somewhere
	QuestTypeCollection

	//QuestTypeEscort is for taking a person somewhere safely, e.g. Barkeep needs drunk patron escorted home
	QuestTypeEscort

	//QuestTypeLocate is for finding an item, e.g. Barkeep lost his wife's jewellery in the forest
	QuestTypeLocate

	//QuestTypeDefend is for defending an item or location, e.g. Barkeep knows town thugs are on their way to collect swag!
	QuestTypeDefend
)

//Quest is a quest containing struct, and contains all parts required to provide a Quest
type Quest struct {
	Type  QuestType
	Steps []QuestStep
}

//String makes the Quest adhere to the Stringer interface
func (q Quest) String() string {
	return "not yet implemented"
}

//QuestStep describe how to solve each stage of a Quest
type QuestStep struct {
	Location     Location
	NPCs         []NPC
	Items        []Item
	Instructions string
}

//GenerateQuest returns a new randomly generated quest of the appropriate QuestType
func GenerateQuest(t QuestType) Quest {
	q := Quest{
		Type: t,
	}

	return q
}
