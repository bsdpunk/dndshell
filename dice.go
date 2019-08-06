package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type DiceRoll struct {
	Number int
	Amount []int
	Type   int
	Total  int
}

func (d DiceRoll) getAmount() (dr DiceRoll) {
	d.Amount = make([]int, d.Number)
	for i := range d.Amount {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(d.Type)
		d.Amount[i] = random + 1
		d.Total = d.Total + d.Amount[i]
		fmt.Println(d)
	}
	return d
}

func main() {

	s := strings.Split(os.Args[1], "d")
	amt, _ := strconv.Atoi(s[0])
	typ, _ := strconv.Atoi(s[1])
	d := DiceRoll{Number: amt, Type: typ}
	d = d.getAmount()
	fmt.Println(d.Total)
}
