package character

import (
	"../commands"
	"./ability"
	class "./class"
	"./race"
	"bufio"
	//	"./skill"
	"../monsters"
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Cl class.Classes
var Rc race.Races

var As ability.AbilityScores
var Ms Monsters.Monsters
var ClassSubs = commands.Commands{
	{
		Name:      "ClassById",
		ShortName: "cid",
		Usage:     "Get Class By Id",
		Category:  "Character",
	},
}
var RaceSubs = commands.Commands{
	{
		Name:      "RaceById",
		ShortName: "rid",
		Usage:     "Get Race By Id",
		Category:  "Character",
	},
}
var AbilitySubs = commands.Commands{
	{
		Name:      "AbilityById",
		ShortName: "aid",
		Usage:     "Get Ability By Id",
		Category:  "Character",
	},
}

var MonsterSubs = commands.Commands{
	{
		Name:      "MonsterById",
		ShortName: "mid",
		Usage:     "Get Monster By Id",
		Category:  "",
	},
}

type Character struct {
	//Player's Name
	PlayerName string
	//Character's Name
	CharacterName string
	//Class
	CharacterClass class.Class
	//Race
	CharacterRace race.Race
	// Ability Scores
	AbilityScores ability.AbilityScores
	//Scores
	CharacterScores []int
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

func Load() {
	Cl.Load()
	Rc.Load()
	Ms.Load()
	As.Load()
}
func (cr *Character) ChooseClass(id string) {
	n, _ := strconv.Atoi(id)
	cr.CharacterClass = Cl.Classes[n]
	return
}

func (cr *Character) ChooseRace(id string) {
	n, _ := strconv.Atoi(id)
	cr.CharacterRace = Rc.Races[n]
	return
}
func (cr *Character) GetScores() {
	var scores []int
	for i := range As.AbilityScores {
		fmt.Println("What is your " + As.AbilityScores[i].Name + " (R for Roll):")
		reader := bufio.NewReader(os.Stdin)
		score, _ := reader.ReadString('\n')
		score = strings.Replace(score, "\n", "", -1)
		n, _ := strconv.Atoi(score)
		scores = append(scores, n)
	}
	cr.CharacterScores = scores
	return
}

func InteractiveCreateCharacter() {
	c := Character{}
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(urlStr)
	fmt.Print("Name$ ")
	Name, _ := reader.ReadString('\n')
	Name = strings.Replace(Name, "\n", "", -1)
	c.PlayerName = Name
	//fmt.Println(c)
	Rc.List()
	fmt.Print("Choose Race: ")
	crace, _ := reader.ReadString('\n')
	crace = strings.Replace(Name, "\n", "", -1)
	c.ChooseRace(crace)
	Cl.List()
	fmt.Print("Choose Class: ")

	cclass, _ := reader.ReadString('\n')
	cclass = strings.Replace(Name, "\n", "", -1)
	c.ChooseRace(cclass)
	c.GetScores()

	e, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	return

}
