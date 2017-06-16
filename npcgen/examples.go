package npcgen

var BanditCaptain = NPC{
	Name: "Bandit Captain",
	BaseStatBlock: StatBlock{
		AbilityScores: AbilityScores{
			Str: 15,
			Dex: 16,
			Con: 14,
			Int: 14,
			Wis: 11,
			Cha: 14,
		},
	},
	ConstantProficiencyModifier: -2,
	HitPoints: DiceFunction{
		Dice:     RepeatDie(DieTypeD8, 10),
		Constant: 0,
	},
	Race: Human,
	Items: []Item{
		Dagger,
		StuddedLeather,
		NecklaceOfFireballs,
		Shield,
	},
}

var Human = RaceTraits{
	Speed: 30,
	Size:  "Medium",
	Type:  "humanoid",

	StatBlockMods: StatBlock{
		AbilityScores: AbilityScores{
			Str: 1,
			Dex: 1,
			Con: 1,
			Int: 1,
			Wis: 1,
			Cha: 1,
		},
	},
	AdultAge: 18,
	MaxAge:   80,
}

var NecklaceOfFireballs = Item{
	Name:       "Necklace of Fireballs (four beads)",
	MinValue:   1800,
	MaxValue:   1200,
	MinWeight:  0.1,
	MaxWeight:  0.3,
	Attributes: "This necklace has 1d6 + 3 beads hanging from it. You can use an action to detach a bead and throw it up to 60 feet away. When it reaches the end of its trajectory, the bead detonates as a 3rd-­‐‐level fireball spell (save DC 15). You can hurl multiple beads, or even the whole necklace, as one action. When you do so, increase the level of the fireball by 1 for each bead beyond the first. ",
	Features: []Feature{
		{
			Name: "Necklace of Fireballs",
			Actions: []Action{
				{
					Name:          "Throw Bead",
					ActionType:    ActionTypeSpell,
					DamageType:    DamageTypeFire,
					AlwaysHits:    true,
					MagicalDamage: true,
					DamageDice:    DiceFunction{Dice: RepeatDie(DieTypeD6, 8)},
					TargetSaveDC:  AbilityScores{Dex: 15},
					MaxRange:      60,
				},
			},
		},
	},
}

var Dagger = Item{
	Name:       "Dagger",
	MinValue:   1,
	MaxValue:   5,
	MinWeight:  1,
	MaxWeight:  1.5,
	Attributes: "",
	Features: []Feature{
		{
			Name: "Dagger",
			Actions: []Action{
				{
					Name:       "Stab",
					ActionType: ActionTypeMeleeWeaponAttack,
					Finesse:    true,
					DamageType: DamageTypePiercing,
					DamageDice: DiceFunction{Dice: []DieType{DieTypeD4}},
				},
				{
					Name:       "Throw",
					ActionType: ActionTypeThrownWeaponAttack,
					Finesse:    true,
					DamageType: DamageTypePiercing,
					DamageDice: DiceFunction{Dice: []DieType{DieTypeD4}},
				},
			},
		},
	},
}

var StuddedLeatherAC = Feature{
	Name: "Studded Leather AC",
	ACSet: &ACSet{
		Base: 12,
		AddMaxAbilityScores: AbilityScores{
			Dex: AbilityScoreUnlimited,
		},
	},
}

var StuddedLeather = Item{
	Name:       "Studded Leather",
	MinValue:   45,
	MaxValue:   45,
	MinWeight:  13,
	MaxWeight:  13,
	Attributes: "",
	Features: []Feature{
		StuddedLeatherAC,
	},
}

var PlusOneStuddedLeather = Item{
	Name:       "Studded Leather",
	MinValue:   45,
	MaxValue:   45,
	MinWeight:  13,
	MaxWeight:  13,
	Attributes: "",
	Features: []Feature{
		StuddedLeatherAC,
		{
			Name:       "Plus One Bonus",
			ACModifier: 1,
		},
	},
}

// TODO: currently we have no way of distinguishing between something that sets AC to 0 and something that doesn't set AC
var Shield = Item{
	Name:       "Shield",
	MinValue:   10,
	MaxValue:   10,
	MinWeight:  6,
	MaxWeight:  6,
	Attributes: "",
	Features: []Feature{
		{
			Name:       "Shield AC",
			ACModifier: 1,
		},
	},
}

// BaseAC is the way we calculate AC for someone not wearing any clothes
var BaseAC = ACSet{
	Base: 10,
	AddMaxAbilityScores: AbilityScores{
		Dex: AbilityScoreUnlimited,
	},
}
