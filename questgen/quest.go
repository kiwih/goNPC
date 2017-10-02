package questgen

//QuestType provides the nomenclature for the type of Quest trope
type QuestType int

type QuestNPC interface {
	GetName() string
}

type QuestLocation interface {
	GetName() string
}

type QuestItem interface {
	GetName() string
}

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

//QuestStep describe how to solve each stage of a Quest, and what happens at each option
type QuestStep struct {
	Instructions string
	Type         QuestType
	Location     QuestLocation
	NPCs         []QuestNPC
	Items        []QuestItem

	NextSteps []NextQuestStep
}

//NextQuestStep
type NextQuestStep struct {
	ChoiceRequirements string
	QuestStep
}

//GenerateQuest returns a new randomly generated quest of the appropriate QuestType
func GenerateQuest(
	possibleLocations []QuestLocation,
	possibleNPCs []QuestNPC,
	possibleItems []QuestItem,
	length int,
) QuestStep {
	q := QuestStep{}

	for i := 0; i < length; i++ {
		//questStep := QuestStep{}

	}

	return q
}
