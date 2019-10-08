package class

import (
	"encoding/json"
	"fmt"
	h "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"strconv"
)

var Home, _ = h.Dir()

type Class struct {
	ClassLevels        ClassLevels          `json:"class_levels"`
	HitDie             float64              `json:"hit_die"`
	ID                 string               `json:"_id"`
	Index              float64              `json:"index"`
	Name               string               `json:"name"`
	Proficiencies      []Proficiencies      `json:"proficiencies"`
	ProficiencyChoices []ProficiencyChoices `json:"proficiency_choices"`
	SavingThrows       []SavingThrows       `json:"saving_throws"`
	StartingEquipment  StartingEquipment    `json:"starting_equipment"`
	Subclasses         []Subclasses         `json:"subclasses"`
	URL                string               `json:"url"`
}
type ClassLevels struct {
	Class string `json:"class"`
	URL   string `json:"url"`
}
type From struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Proficiencies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type ProficiencyChoices struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type SavingThrows struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingEquipment struct {
	Class string `json:"class"`
	URL   string `json:"url"`
}
type Subclasses struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Classes struct {
	Classes []Class `json:"classes"`
}

func (cl *Classes) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/class.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &cl.Classes); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (cl *Classes) ClassById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(cl.Classes[n]) //.Strength)
	return
}

func (cl *Classes) List() {
	for i := range cl.Classes {
		fmt.Print((cl.Classes[i].Index - 1))
		fmt.Println(" " + cl.Classes[i].Name)
	}
	//fmt.Println(cl.Classes)
	return

}
