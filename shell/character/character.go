package character

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bsdpunk/dndshell/shell/character/ability"
	class "github.com/bsdpunk/dndshell/shell/character/class"
	"github.com/bsdpunk/dndshell/shell/character/levels"
	"github.com/bsdpunk/dndshell/shell/character/race"
	"github.com/bsdpunk/dndshell/shell/character/spells"
	"github.com/bsdpunk/dndshell/shell/commands"
	"github.com/bsdpunk/dndshell/shell/monsters"
	"github.com/gobs/readline"
	"os"
	"strconv"
	"strings"
)

var Cl class.Classes
var Rc race.Races
var Ls levels.Levels
var Sp spells.Spells
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
var SpellSubs = commands.Commands{
	{
		Name:      "SpellsById",
		ShortName: "sid",
		Usage:     "Get Spell By Id",
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
	{
		Name:      "MonsterByCR",
		ShortName: "mbc",
		Usage:     "Get Monster By CR",
		Category:  "",
	},
	{
		Name:      "MonsterByGuess",
		ShortName: "mbg",
		Usage:     "Get Monster By Guess",
		Category:  "",
	},
	{
		Name:      "MonsterByName",
		ShortName: "mbn",
		Usage:     "Get Monster By Name",
		Category:  "",
	},
	{
		Name:      "SBById",
		ShortName: "sbid",
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

var levToP = map[int]int{
	1:  2,
	2:  2,
	3:  2,
	4:  2,
	5:  3,
	6:  3,
	7:  3,
	8:  3,
	9:  4,
	10: 4,
	11: 4,
	12: 4,
	13: 5,
	14: 5,
	15: 5,
	16: 5,
	17: 6,
	18: 6,
	19: 6,
	20: 6,
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
	//Modifiers
	CharacterModifiers []int
	//Actions
	Actions     int
	BonusAction int
	//Level Prof Bonus Map
	Prof int
	//Health
	HitPoints int
	//Armor
	ArmorClass int
	//Initiative
	Initiative int
	//Level
	Level int
}

func Load() {
	Cl.Load()
	Rc.Load()
	Ms.Load()
	As.Load()
	Ls.Load()
	Sp.Load()
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
	var mods []int
	for i := range As.AbilityScores {
		fmt.Println("What is your " + As.AbilityScores[i].Name + " (R for Roll):")
		reader := bufio.NewReader(os.Stdin)
		score, _ := reader.ReadString('\n')
		score = strings.Replace(score, "\n", "", -1)
		n, _ := strconv.Atoi(score)
		if n == 12 || n == 13 {
			mods = append(mods, 1)
		} else if n == 8 || n == 9 {
			mods = append(mods, -1)
		} else {
			mods = append(mods, 0)
		}
		scores = append(scores, n)
	}
	cr.CharacterScores = scores
	cr.CharacterModifiers = mods

	return
}
func (cr *Character) GetNPCScores() {
	var scores []string
	var mods []int
	for i := range As.AbilityScores {
		fmt.Println("What is your " + As.AbilityScores[i].Name + " (B for Beefy, A for Average, W for Weak):")
		reader := bufio.NewReader(os.Stdin)
		score, _ := reader.ReadString('\n')
		score = strings.Replace(score, "\n", "", -1)
		//n, _ := strconv.Atoi(score)
		if score == "b" || score == "B" {
			mods = append(mods, 3)
		} else if score == "a" || score == "A" {
			mods = append(mods, 0)
		} else {
			mods = append(mods, -1)
		}
		scores = append(scores, score)
	}
	//cr.CharacterScores = scores
	cr.CharacterModifiers = mods

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

	promptL := "What Level: "
	readL := readline.ReadLine(&promptL)
	c.Level, _ = strconv.Atoi(*readL)
	c.Prof = levToP[c.Level]
	c.GetScores()

	e, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	return

}

func InteractiveCreateNPC() {
	c := Character{}
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(urlStr)
	fmt.Print("Name: ")
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

	promptL := "What Level: "
	readL := readline.ReadLine(&promptL)
	c.Level, _ = strconv.Atoi(*readL)
	c.Prof = levToP[c.Level]
	c.GetNPCScores()

	e, err := json.Marshal(&c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	return

}
