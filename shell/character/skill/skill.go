package skill

import (
	h "github.com/mitchellh/go-homedir"
)

var Home, _ = h.Dir()

type Skill struct {
	AbilityScore AbilityScore `json:"ability_score"`
	Desc         []string     `json:"desc"`
	Index        float64      `json:"index"`
	Name         string       `json:"name"`
	Skills       []Skills     `json:"Skills"`
	URL          string       `json:"url"`
}
type AbilityScore struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Skills struct {
	Skills []Skill
}

func (rc *Skills) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/skills.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.Skills); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *Skills) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.Skills[n]) //.Strength)
	return
}

func (rc *Skills) List() {
	for i := range rc.Skills {
		fmt.Print((rc.Skills[i].Index - 1))
		fmt.Println(" " + rc.Skills[i].Name)
	}
	//fmt.Println(cl.Skills)
	return

}
