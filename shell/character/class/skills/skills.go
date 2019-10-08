package skills

type Skill struct {
	Index        int      `json:"index"`
	Name         string   `json:"name"`
	Desc         []string `json:"desc"`
	AbilityScore struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"ability_score"`
	URL string `json:"url"`
}

type Skills struct {
	Skills []Skill
}
