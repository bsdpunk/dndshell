package character

import (
	"../commands"
	"./ability"
	class "./class"
	"./levels"
	"./race"
	"bufio"
	"github.com/gobs/readline"
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
var Ls levels.Levels

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
var LevelSubs = commands.Commands{

	{
		Name:      "LevelById",
		ShortName: "lid",
		Usage:     "Get Level By Id",
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
	// Ability Score Definitions and Relationships
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
	Ls.Load()
}
func (cr *Character) ChooseClass(id *string) {
	idt := *id
	n, _ := strconv.Atoi(idt)
	cr.CharacterClass = Cl.Classes[n]
	fmt.Println(Cl.Classes[n])
	return
}

func (cr *Character) ChooseRace(id *string) {
	idt := *id
	n, _ := strconv.Atoi(idt)
	fmt.Println(id)
	fmt.Println(n)
	cr.CharacterRace = Rc.Races[n]
	fmt.Println(Rc.Races[n])
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
	promptN := "Character Name: "
	readN := readline.ReadLine(&promptN)
	cn := *readN
	c.CharacterName = cn
	//fmt.Println(c)
	Rc.List()
	promptR := "Choose Race: "
	readR := readline.ReadLine(&promptR)

	c.ChooseRace(readR)
	Cl.List()
	promptC := "Choose Class: "
	readC := readline.ReadLine(&promptC)

	c.ChooseClass(readC)
	c.GetScores()

	e, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	return

}
