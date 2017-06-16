package npcgen

//An ActionType is used for identifying the kind of action this is
type ActionType string

const (
	//ActionTypeMeleeAttack is for melee actions e.g. hit with a sword (uses STR)
	ActionTypeMeleeAttack ActionType = "Melee Attack"

	//ActionTypeMeleeFinesseAttack is for melee finesse attacks, e.g. hit with a scimitar (uses better of STR or DEX)
	ActionTypeMeleeFinesseAttack ActionType = "Melee (Finesse) Attack"

	//ActionTypeRangedAttack is for ranged actions e.g. i shoot with a bow (uses DEX)
	ActionTypeRangedAttack ActionType = "Ranged Attack"

	//ActionTypeMeleeSpellAttack is for melee spell attack actions that use spell modifier for attack roll e.g. shocking grasp (uses SPELL MOD)
	ActionTypeMeleeSpellAttack ActionType = "Melee (Spell) Attack"

	//ActionTypeRangedSpellAttack is for ranged spell attacks that use spell modifier for attack roll e.g. ray of frost (uses SPELL MOD)
	ActionTypeRangedSpellAttack ActionType = "Ranged (Spell) Attack"

	//ActionTypeThrownAttack is for attack rolls that are thrown, e.g. I throw the javelin (uses STR)
	ActionTypeThrownAttack ActionType = "Thrown Attack"

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

	DamageDice DiceFunction //if appropriate

	UseSpellSaveDC bool          //If true, replaces the positive values in TargetSaveDC with the attackers Spell Save DC
	TargetSaveDC   AbilityScores //If the target needs to make a save, this stores the save value (use any positive value in the appropriate slot if it's based on caster spell save dc)

	DamageType    DamageType
	MagicalDamage bool

	ShortRange int //if set, the time which separates an attack from an attack with disadvantage
	MaxRange   int //the maximum range for an attack, if ShortRange is set, this is the range that attacks will have disadv, if not set, this is the max range
}

//if Melee and short/longRange, it's probably a melee/thrown weapon
//if not Melee and short/longRange, it's probably a ranged weapon

//AttackModifier calculates the attack modifier for an action given the attacking npc "n"
func (a Action) AttackModifier(n NPC) int {
	//ActionTypeSpell has no modifier
	if a.ActionType == ActionTypeSpell || a.ActionType == ActionTypeRitual {
		return 0
	}

	//Actions need to depend on ability scores, e.g. when casting a spell we use spellcasting ability, etc, thrown weapons vs other things etc

	AttackStr := n.StrAttackModifier() + n.ProficiencyBonus //we assume that NPCs don't ever attack with things they aren't proficient with
	AttackDex := n.DexAttackModifier() + n.ProficiencyBonus
	AttackSpell := n.SpellAttackModifier() + n.ProficiencyBonus

	//TODO: Based on the player statblock, and the type of action, calculate the attack modifier

	switch a.ActionType {
	case ActionTypeMeleeAttack:
		return AttackStr
	case ActionTypeMeleeFinesseAttack:
		if AttackStr > AttackDex {
			return AttackStr
		}
		return AttackDex
	case ActionTypeRangedAttack:
		return AttackDex
	case ActionTypeMeleeSpellAttack:
		return AttackSpell
	case ActionTypeRangedSpellAttack:
		return AttackSpell
	case ActionTypeThrownAttack:
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
