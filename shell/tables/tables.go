package Monsters

import (
	"encoding/json"
	"fmt"
	"github.com/elgs/jsonql"
	h "github.com/mitchellh/go-homedir"
	"github.com/tatsushid/go-prettytable"
	"io/ioutil"
	"os"
	"strconv"
)

var Home, _ = h.Dir()

type Table struct {
	Name  string
	table map[int]string
}

type Tables struct {
	Tables     []Table
	jsonString string
}

func (ms *Tables) List() {

	for i := range ms.Tables {
		fmt.Print((ms.Tables[i].Index - 1))
		fmt.Println(" " + ms.Tables[i].Name)
	}
	return

}

func (ms *Tables) TableById(id string) {
	n, _ := strconv.Atoi(id)
	fmt.Println(ms.Tables[n]) //.Strength)
	return
}

func (ms *Tables) ByName(name string) {
	parser, err := jsonql.NewStringQuery(ms.jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("name='" + name + "'"))
}

func (ms *Tables) ByGuess(name string) {
	parser, err := jsonql.NewStringQuery(ms.jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("name~='" + name + "'"))
}

func (ms *Tables) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/monsters.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	ms.jsonString = string(byteValue)
	if err := json.Unmarshal(byteValue, &ms.Tables); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}
