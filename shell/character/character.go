package main

import "fmt"
import "./class"
import "./race"

var Class = class.Class

type Character struct {
	//Player's Name
	PlayerName string
	//Characther's Name
	CharacterName string
	//Class
	CharacterClass Class
	//Race
	CharacerRace race.Race
	// Ability Scores
	AbilityScores AS
	//Actions
	Actions     int
	BonusAction int
	//Health
	HitPoints int
	//Armor
	ArmorClass int
}

type AS struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Wisdom       int
	Constitution int
	Charisma     int
}

func main() {
	fmt.Println("vim-go")
}
