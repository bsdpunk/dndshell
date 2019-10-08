package race

import (
	"encoding/json"
	"fmt"
	h "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"strconv"
)

var Home, _ = h.Dir()

type Race struct {
	AbilityBonusOptions        AbilityBonusOptions        `json:"ability_bonus_options"`
	AbilityBonuses             []AbilityBonuses           `json:"ability_bonuses"`
	Age                        string                     `json:"age"`
	Alignment                  string                     `json:"alignment"`
	Index                      float64                    `json:"index"`
	LanguageDesc               string                     `json:"language_desc"`
	LanguageOptions            LanguageOptions            `json:"language_options"`
	Languages                  []Languages                `json:"languages"`
	Name                       string                     `json:"name"`
	Races                      []Races                    `json:"Races"`
	Size                       string                     `json:"size"`
	SizeDescription            string                     `json:"size_description"`
	Speed                      float64                    `json:"speed"`
	StartingProficiencies      []StartingProficiencies    `json:"starting_proficiencies"`
	StartingProficiencyOptions StartingProficiencyOptions `json:"starting_proficiency_options"`
	Subraces                   []Subraces                 `json:"subraces"`
	TraitOptions               TraitOptions               `json:"trait_options"`
	Traits                     []Traits                   `json:"traits"`
	URL                        string                     `json:"url"`
}
type AbilityBonusOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type AbilityBonuses struct {
	Bonus float64 `json:"bonus"`
	Name  string  `json:"name"`
	URL   string  `json:"url"`
}
type From struct {
	Bonus float64 `json:"bonus"`
	Name  string  `json:"name"`
	URL   string  `json:"url"`
}
type LanguageOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Languages struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingProficiencies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingProficiencyOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Subraces struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type TraitOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Traits struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Races struct {
	Races []Race
}

func (rc *Races) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/race.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.Races); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *Races) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.Races[n]) //.Strength)
	return
}

func (rc *Races) List() {
	for i := range rc.Races {
		fmt.Print((rc.Races[i].Index - 1))
		fmt.Println(" " + rc.Races[i].Name)
	}
	//fmt.Println(cl.Races)
	return

}
