package npcgen

var BanditCaptain = NPC{
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
	HitPoints: DiceFunction{
		Dice:     RepeatDie(DieTypeD8, 10),
		Constant: 0,
	},
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
			Name: "Fireballs",
			Actions: []Action{
				{
					Name:         "Throw Bead",
					ActionType:   ActionTypeSpell,
					DamageType:   DamageTypeFire,
					DamageDice:   DiceFunction{Dice: RepeatDie(DieTypeD6, 8)},
					TargetSaveDC: AbilityScores{Dex: 15},
					MaxRange:     60,
				},
			},
		},
	},
}
