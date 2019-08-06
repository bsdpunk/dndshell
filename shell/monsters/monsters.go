package monsters

type Monster struct {
	Actions               []Actions          `json:"actions"`
	Alignment             string             `json:"alignment"`
	ArmorClass            float64            `json:"armor_class"`
	ChallengeRating       float64            `json:"challenge_rating"`
	Charisma              float64            `json:"charisma"`
	CharismaSave          float64            `json:"charisma_save"`
	ConditionImmunities   string             `json:"condition_immunities"`
	Constitution          float64            `json:"constitution"`
	ConstitutionSave      float64            `json:"constitution_save"`
	DamageImmunities      string             `json:"damage_immunities"`
	DamageResistances     string             `json:"damage_resistances"`
	DamageVulnerabilities string             `json:"damage_vulnerabilities"`
	Dexterity             float64            `json:"dexterity"`
	DexteritySave         float64            `json:"dexterity_save"`
	HitDice               string             `json:"hit_dice"`
	HitPoints             float64            `json:"hit_points"`
	ID                    string             `json:"_id"`
	Index                 float64            `json:"index"`
	Intelligence          float64            `json:"intelligence"`
	Languages             string             `json:"languages"`
	LegendaryActions      []LegendaryActions `json:"legendary_actions"`
	Name                  string             `json:"name"`
	Perception            float64            `json:"perception"`
	Senses                string             `json:"senses"`
	Size                  string             `json:"size"`
	SpecialAbilities      []SpecialAbilities `json:"special_abilities"`
	Speed                 string             `json:"speed"`
	Stealth               float64            `json:"stealth"`
	Strength              float64            `json:"strength"`
	Subtype               string             `json:"subtype"`
	Type                  string             `json:"type"`
	URL                   string             `json:"url"`
	Wisdom                float64            `json:"wisdom"`
	WisdomSave            float64            `json:"wisdom_save"`
}
type Actions struct {
	AttackBonus float64 `json:"attack_bonus"`
	DamageBonus float64 `json:"damage_bonus"`
	DamageDice  string  `json:"damage_dice"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}
type LegendaryActions struct {
	AttackBonus float64 `json:"attack_bonus"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}

type Monsters struct {
	Monsters []Monster `json:"monsters"`
	//Monster  Monster    `json:"monster"`
}

type Monster struct {
	Name        string `json:"name"`
	Id          int    `json:"id"`
	Age         string `json:"HITPOINTS"`
	AC          int    `json:"AC"`
	CR          int    `json:"CR"`
	XP          string `json:"XP"`
	Description string `json:"Descriptions"`
	//squad
}
