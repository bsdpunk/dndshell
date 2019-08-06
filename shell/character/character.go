package character

import "fmt"
import "./class"
import "./race"
import "./inv"
import "bufio"
import "os"
import "strings"
import "encoding/json"

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
	AbilityScores AS
	// Ability Mods
	AbilityMods Mods
	//Actions
	Actions     int
	BonusAction int
	//Health
	HitPoints int
	//Armor
	ArmorClass int
	//Inventory
	Inventory inv.Inventory
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

type AS struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Wisdom       int
	Constitution int
	Charisma     int
}
type Mods struct {
	StrengthMod     int
	DexterityMod    int
	IntelligenceMod int
	WisdomMod       int
	ConstitutionMod int
	CharismaMod     int
}
