package spells

import (
	"encoding/json"
	"fmt"
	"github.com/tatsushid/go-prettytable"
	"io/ioutil"
	//"log"
	"os"
	//"os/exec"
	h "github.com/mitchellh/go-homedir"
	"strconv"
	"strings"
)

var Home, _ = h.Dir()

type Spell struct {
	Index         int      `json:"index"`
	Name          string   `json:"name"`
	Desc          []string `json:"desc"`
	HigherLevel   []string `json:"higher_level,omitempty"`
	Page          string   `json:"page"`
	Range         string   `json:"range"`
	Components    []string `json:"components"`
	Material      string   `json:"material,omitempty"`
	Ritual        bool     `json:"ritual"`
	Duration      string   `json:"duration"`
	Concentration bool     `json:"concentration"`
	CastingTime   string   `json:"casting_time"`
	Level         int      `json:"level"`
	School        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"school"`
	Classes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"classes"`
	Subclasses []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"subclasses"`
	URL string `json:"url"`
}

type Spells struct {
	Spells []Spell
}

func (rc *Spells) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/spells.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.Spells); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *Spells) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.Spells[n]) //.Strength)
	return
}

func (rc *Spells) List() {
	for i := range rc.Spells {
		fmt.Print((rc.Spells[i].Index - 1))
		fmt.Println(" " + rc.Spells[i].Name)
	}
	//fmt.Println(cl.Spells)
	return

}

func (sp *Spells) SpellBook() {
	//cmd := exec.Command("stty", "size")
	//cmd.Stdin = os.Stdin
	//out, err := cmd.Output()
	//	fmt.Printf("out: %#v\n", string(out))
	//	fmt.Printf("err: %#v\n", err)
	//if err != nil {
	//	log.Fatal(err)
	//}
	i := 0
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "Left", MinWidth: 80, MaxWidth: 150},
		{Header: "Right", MinWidth: 80, MaxWidth: 150},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	tbl.AddRow(strings.Join(sp.Spells[i].Desc, " "), strings.Join(sp.Spells[i+1].Desc, " "))
	tbl.AddRow(strings.Join(sp.Spells[i+2].Desc, " "), strings.Join(sp.Spells[i+3].Desc, " "))
	tbl.AddRow(strings.Join(sp.Spells[i+5].Desc, " "), strings.Join(sp.Spells[i+6].Desc, " "))
	//tbl.AddRow(sp.Spells[i+6].Desc, sp.Spells[i+7].Desc, sp.Spells[i+8].Desc)
	tbl.Print()
	//fmt.Println(err)

}
