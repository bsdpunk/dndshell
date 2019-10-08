package ability

import (
	"encoding/json"
	"fmt"
	h "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"strconv"
)

var Home, _ = h.Dir()

type AbilityScore struct {
	AbilityScores []AbilityScores `json:"Ability-Scores"`
	Desc          []string        `json:"desc"`
	FullName      string          `json:"full_name"`
	Index         float64         `json:"index"`
	Name          string          `json:"name"`
	Skills        []Skills        `json:"skills"`
	URL           string          `json:"url"`
}
type Skills struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AbilityScores struct {
	AbilityScores []AbilityScore
}

func (rc *AbilityScores) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/ability.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.AbilityScores); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *AbilityScores) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.AbilityScores[n]) //.Strength)
	return
}

func (rc *AbilityScores) List() {
	for i := range rc.AbilityScores {
		fmt.Print((rc.AbilityScores[i].Index - 1))
		fmt.Println(" " + rc.AbilityScores[i].Name)
	}
	//fmt.Println(cl.AbilityScores)
	return

}
