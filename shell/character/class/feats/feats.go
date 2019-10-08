package feats

import "fmt"

type Feat struct {
	Index int `json:"index"`
	Class struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"class"`
	Name     string   `json:"name"`
	Level    int      `json:"level,omitempty"`
	Desc     []string `json:"desc"`
	URL      string   `json:"url"`
	Subclass struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"subclass,omitempty"`
}

type Feats struct {
	Feats []Feat
}

func (rc *Feats) Load() {
	jsonFile, err := os.Open("./json/feats.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &rc.Feats); err != nil {
		fmt.Println(err)
	}
	//json.Unmarshal(byteValue, &ms)
	return
}

func (rc *Feats) ById(id string) {

	n, _ := strconv.Atoi(id)
	fmt.Println(rc.Feats[n]) //.Strength)
	return
}

func (rc *Feats) List() {
	for i := range rc.Feats {
		fmt.Print((rc.Feats[i].Index - 1))
		fmt.Println(" " + rc.Feats[i].Name)
	}
	//fmt.Println(cl.Feats)
	return

}
