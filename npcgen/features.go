package npcgen

//A Feature is like a passive or active trait and can do pretty much anything, including actions etc
// (they are applied to items, races, etc)
type Feature struct {
	Name                  string
	Description           string
	AbilityScoreModifiers AbilityScores
	ACSet                 *ACSet
	ACModifier            int      //Adds this to character's AC no questions asked
	Actions               []Action //e.g. you can stab with a dagger
	Reactions             []Action //e.g. you can parry with a sword
}
