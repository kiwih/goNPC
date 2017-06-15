package npcgen

//ItemType describes the type of item an item is (e.g. Jewellery, furniture, clothing)
type ItemType int

const (
	//ItemTypeJewellery is jewellery
	ItemTypeJewellery ItemType = iota
)

//An Item can be any carried or worn item, such as armour, jewellery, etc
type Item struct {
	Type ItemType
	Name string

	MinValue  int
	MaxValue  int
	MinWeight int
	MaxWeight int

	Attributes string
	Features   []Feature //e.g. an item might increase your strength //e.g. you can stab with a dagger //e.g. you can parry with a sword
}

//RandomItem returns a random item
func RandomItem(t ItemType) Item {
	return Item{}
}
