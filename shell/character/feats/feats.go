package feats

import (
	h "github.com/mitchellh/go-homedir"
)

var Home, _ = h.Dir()

type Feats struct {
	Class Class    `json:"class"`
	Desc  []string `json:"desc"`
	ID    string   `json:"_id"`
	Index float64  `json:"index"`
	Level float64  `json:"level"`
	Name  string   `json:"name"`
	URL   string   `json:"url"`
}
type Class struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (ms *Feats) FeatById(id string) {
	n, _ := strconv.Atoi(id)
	fmt.Println(ms.Feats[n]) //.Strength)
	return
}

func (ms *Feats) Load() {
	jsonFile, err := os.Open(Home + "/go/src/github.com/bsdpunk/dndshell/json/monsters.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &ms.Feats); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}
