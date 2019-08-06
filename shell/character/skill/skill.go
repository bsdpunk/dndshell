package skill

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
