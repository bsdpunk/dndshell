package character

import (
	"./ability"
	"./class"
	"./race"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Character struct {
	//Player's Name
	PlayerName string
	//Characther's Name
	CharacterName string
	//Class
	CharacterClass class.Classes
	//Race
	CharacerRace race.Race
	// Ability Scores
	AbilityScores ability.AbilityScores
	//Actions
	Actions     int
	BonusAction int
	//Health
	HitPoints int
	//Armor
	ArmorClass int
	//Initiative
	Initiative int
	//
}

func InteractiveCreateCharacter() {
	c := Character{}
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(urlStr)
	fmt.Print("Name$ ")
	Name, _ := reader.ReadString('\n')
	Name = strings.Replace(Name, "\n", "", -1)
	c.PlayerName = Name
	fmt.Println(c)
	e, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	return

}

func LoadClasses() (cs class.Classes) {
	jsonFile, err := os.Open("./json/classes.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cs)
	return cs
}
