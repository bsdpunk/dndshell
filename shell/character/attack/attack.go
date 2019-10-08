package attack

import "../../dice"

//import "fmt"

//func main() {
//	fmt.Println("vim-go")
//}

type Attack struct {
	AttackBonus float64 `json:"attack_bonus"`
	DamageBonus int     `json:"damage_bonus"`
	DamageDice  string  `json:"damage_dice"`
	Desc        string  `json:"desc"`
	Name        string  `json:"name"`
}

type Attacks struct {
	Attacks []Attack
}

func (a Attack) AttackRoll(ability int, prof int, other int) int {
	return dice.GetAmount("1d20") + ability + prof + other
}
