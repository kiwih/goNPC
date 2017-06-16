package npcgen

import "fmt"

//An ActionType is used for identifying the kind of action this is
type ActionType string

const (
	//ActionTypeMeleeWeaponAttack is for melee actions e.g. hit with a sword (uses STR)
	ActionTypeMeleeWeaponAttack ActionType = "Melee Weapon Attack"

	//ActionTypeRangedWeaponAttack is for ranged actions e.g. i shoot with a bow (uses DEX)
	ActionTypeRangedWeaponAttack ActionType = "Ranged Weapon Attack"

	//ActionTypeMeleeSpellAttack is for melee spell attack actions that use spell modifier for attack roll e.g. shocking grasp (uses SPELL MOD)
	ActionTypeMeleeSpellAttack ActionType = "Melee Spell Attack"

	//ActionTypeRangedSpellAttack is for ranged spell attacks that use spell modifier for attack roll e.g. ray of frost (uses SPELL MOD)
	ActionTypeRangedSpellAttack ActionType = "Ranged Spell Attack"

	//ActionTypeThrownWeaponAttack is for attack rolls that are thrown, e.g. I throw the javelin (uses STR)
	ActionTypeThrownWeaponAttack ActionType = "Thrown Weapon Attack"

	//ActionTypeSpell is for spell actions that aren't attacks, or are AOE, e.g. I use dissonant whispers / I throw a fireball (uses SPELL SAVE DC, if appropriate)
	ActionTypeSpell ActionType = "Spell"

	//ActionTypeRitual is for casting ritual spells (uses SPELL SAVE DC, if appropriate)
	ActionTypeRitual ActionType = "Ritual"
)

//An Action is something that an NPC/PC can do when it has the appropriate thing that grants this action
//e.g. a Sword would give the sword action
type Action struct {
	Name       string
	ActionType ActionType //we can use this to determine the attack roll modifier
	Finesse    bool

	DamageDice DiceFunction //if appropriate

	AlwaysHits bool //this should be set if it is something that either always hits, or something that provides only a "save for half" e.g. fireballs

	UseSpellSaveDC bool          //If true, replaces the positive values in TargetSaveDC with the attackers Spell Save DC
	TargetSaveDC   AbilityScores //If the target needs to make a save, this stores the save value (use any positive value in the appropriate slot if it's based on caster spell save dc)

	DamageType    DamageType
	MagicalDamage bool

	Reach      bool
	ShortRange int //if set, the time which separates an attack from an attack with disadvantage
	MaxRange   int //the maximum range for an attack, if ShortRange is set, this is the range that attacks will have disadv, if not set, this is the max range
}

//AttackModifier calculates the attack modifier for an action given the attacking npc "n"
func (a Action) AttackModifier(n NPC) int {
	return a.attackVal(n, true)
}

//attackVal calculates the attack / damage modifiers for an action given the attacking npc "n" and whether or not to add the proficiency bonus
func (a Action) attackVal(n NPC, addProf bool) int {
	//ActionTypeSpell has no modifier
	if a.ActionType == ActionTypeSpell || a.ActionType == ActionTypeRitual {
		return 0
	}

	//Actions need to depend on ability scores, e.g. when casting a spell we use spellcasting ability, etc, thrown weapons vs other things etc

	AttackStr := n.StrAttackModifier(addProf) //we assume that NPCs don't ever attack with things they aren't proficient with
	AttackDex := n.DexAttackModifier(addProf)
	AttackSpell := n.SpellAttackModifier(addProf)

	//Based on the player stats, and the type of action, calculate the attack modifier

	switch a.ActionType {
	case ActionTypeMeleeWeaponAttack:
		if a.Finesse && AttackDex > AttackStr {
			return AttackDex
		}
		return AttackStr
	case ActionTypeRangedWeaponAttack:
		if a.Finesse && AttackStr > AttackDex {
			return AttackStr
		}
		return AttackDex
	case ActionTypeMeleeSpellAttack:
		return AttackSpell
	case ActionTypeRangedSpellAttack:
		return AttackSpell
	case ActionTypeThrownWeaponAttack:
		if a.Finesse && AttackDex > AttackStr {
			return AttackDex
		}
		return AttackStr

	default: //shouldn't occur unless ActionType is invalid
		return 0
	}
}

//SaveDC gets the SaveDC for the target given the caster/attacking npc "n"
func (a Action) SaveDC(n NPC) AbilityScores {
	if a.UseSpellSaveDC {
		s := a.TargetSaveDC
		if s.Cha > 0 {
			s.Cha = AbilityScore(n.SpellSaveDC())
		}
		if s.Con > 0 {
			s.Con = AbilityScore(n.SpellSaveDC())
		}
		if s.Dex > 0 {
			s.Dex = AbilityScore(n.SpellSaveDC())
		}
		if s.Int > 0 {
			s.Int = AbilityScore(n.SpellSaveDC())
		}
		if s.Str > 0 {
			s.Str = AbilityScore(n.SpellSaveDC())
		}
		if s.Wis > 0 {
			s.Wis = AbilityScore(n.SpellSaveDC())
		}
		return s
	}
	return a.TargetSaveDC
}

//RangeString formats the range information for the action into something easy to read
func (a Action) RangeString() string {
	if a.ShortRange != 0 {
		return fmt.Sprintf("range %d/%d ft.", a.ShortRange, a.MaxRange)
	}
	if a.MaxRange != 0 {
		return fmt.Sprintf("range %d ft.", a.MaxRange)
	}
	if a.Reach {
		return "reach 10 ft."
	}
	return "reach 5 ft."
}

//Damage returns the DamageDice function taking into account all player stats
func (a Action) Damage(n NPC) DiceFunction {
	d := a.DamageDice
	d.Constant += a.attackVal(n, false)
	return d
}

//DamageString returns the DamageDice function as a nice string, taking into account all player stats
func (a Action) DamageString(n NPC) string {
	return a.Damage(n).String()
}

//AttackString returns the nice "+x to hit" string
// it can also say "DC x Dex to half"
func (a Action) AttackString(n NPC) string {
	s := ""
	if !a.AlwaysHits {
		s += fmt.Sprintf("+%d to hit", a.AttackModifier(n))
	}
	saves := a.SaveDC(n)
	saveStr := ""
	if saves.Cha > 0 {
		saveStr += fmt.Sprintf("DC %d Charisma to half", saves.Cha)
	}
	if saves.Con > 0 {
		saveStr += fmt.Sprintf("DC %d Constitution to half", saves.Con)
	}
	if saves.Dex > 0 {
		saveStr += fmt.Sprintf("DC %d Dexterity to half", saves.Dex)
	}
	if saves.Int > 0 {
		saveStr += fmt.Sprintf("DC %d Intelligence to half", saves.Int)
	}
	if saves.Str > 0 {
		saveStr += fmt.Sprintf("DC %d Strength to half", saves.Str)
	}
	if saves.Wis > 0 {
		saveStr += fmt.Sprintf("DC %d Wisdom to half", saves.Wis)
	}
	if saveStr != "" && s != "" {
		return s + "(" + saveStr + ")"
	}
	if saveStr != "" {
		return saveStr
	}
	return s
}

//DamageTypeString gets a nice "magical fire" string for damage type on an action
func (a Action) DamageTypeString() string {
	s := ""
	if a.MagicalDamage {
		s += "magical "
	}
	return s + string(a.DamageType)
}
