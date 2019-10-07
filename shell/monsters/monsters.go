package Monsters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Monster struct {
	Acrobatics            float64            `json:"acrobatics"`
	Actions               []Actions          `json:"actions"`
	Alignment             string             `json:"alignment"`
	Arcana                float64            `json:"arcana"`
	ArmorClass            float64            `json:"armor_class"`
	Athletics             float64            `json:"athletics"`
	ChallengeRating       float64            `json:"challenge_rating"`
	Charisma              float64            `json:"charisma"`
	CharismaSave          float64            `json:"charisma_save"`
	ConditionImmunities   []string           `json:"condition_immunities"`
	Constitution          float64            `json:"constitution"`
	ConstitutionSave      float64            `json:"constitution_save"`
	DamageImmunities      []string           `json:"damage_immunities"`
	DamageResistances     []string           `json:"damage_resistances"`
	DamageVulnerabilities []string           `json:"damage_vulnerabilities"`
	Deception             float64            `json:"deception"`
	Dexterity             float64            `json:"dexterity"`
	DexteritySave         float64            `json:"dexterity_save"`
	History               float64            `json:"history"`
	HitDice               string             `json:"hit_dice"`
	HitPoints             float64            `json:"hit_points"`
	Index                 float64            `json:"index"`
	Insight               float64            `json:"insight"`
	Intelligence          float64            `json:"intelligence"`
	IntelligenceSave      float64            `json:"intelligence_save"`
	Intimidation          float64            `json:"intimidation"`
	Investigation         float64            `json:"investigation"`
	Languages             string             `json:"languages"`
	LegendaryActions      []LegendaryActions `json:"legendary_actions"`
	Medicine              float64            `json:"medicine"`
	Monsters              []Monster          `json:"Monsters"`
	Name                  string             `json:"name"`
	Nature                float64            `json:"nature"`
	OtherSpeeds           []OtherSpeeds      `json:"other_speeds"`
	Perception            float64            `json:"perception"`
	Performance           float64            `json:"performance"`
	Persuasion            float64            `json:"persuasion"`
	Reactions             []Reactions        `json:"reactions"`
	Religion              float64            `json:"religion"`
	Senses                string             `json:"senses"`
	Size                  string             `json:"size"`
	SpecialAbilities      []SpecialAbilities `json:"special_abilities"`
	Speed                 Speed              `json:"speed"`
	Stealth               float64            `json:"stealth"`
	Strength              float64            `json:"strength"`
	StrengthSave          float64            `json:"strength_save"`
	Subtype               string             `json:"subtype"`
	Survival              float64            `json:"survival"`
	Type                  string             `json:"type"`
	URL                   string             `json:"url"`
	Wisdom                float64            `json:"wisdom"`
	WisdomSave            float64            `json:"wisdom_save"`
}
type Actions struct {
	AttackBonus float64 `json:"attack_bonus"`
	//DamageBonus float64    `json:"damage_bonus"`
	DamageDice string `json:"damage_dice"`
	Desc       string `json:"desc"`
	Name       string `json:"name"`
}
type LegendaryActions struct {
	AttackBonus float64 `json:"attack_bonus"`
	DamageDice  string  `json:"damage_dice"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}
type OtherSpeeds struct {
	Form  string `json:"form"`
	Speed Speed  `json:"speed"`
}
type Reactions struct {
	AttackBonus float64 `json:"attack_bonus"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}
type SpecialAbilities struct {
	AttackBonus float64 `json:"attack_bonus"`
	DamageDice  string  `json:"damage_dice"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}
type Speed struct {
	Burrow string `json:"burrow"`
	Climb  string `json:"climb"`
	Fly    string `json:"fly"`
	Hover  bool   `json:"hover"`
	Swim   string `json:"swim"`
	Walk   string `json:"walk"`
}
type Monsters struct {
	Monsters []Monster
}

func (ms *Monsters) List() {

	for i := range ms.Monsters {
		fmt.Print((ms.Monsters[i].Index - 1))
		fmt.Println(" " + ms.Monsters[i].Name)
	}
	return

}

func (ms *Monsters) MonsterById(id string) {
	n, _ := strconv.Atoi(id)
	fmt.Println(ms.Monsters[n]) //.Strength)
	return
}

func (ms *Monsters) Load() {
	jsonFile, err := os.Open("./json/monsters.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &ms.Monsters); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}
