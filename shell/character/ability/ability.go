package ability

type AbilityScores struct {
	AbilityScores []AbilityScores `json:"Ability-Scores"`
	Desc          []string        `json:"desc"`
	FullName      string          `json:"full_name"`
	Index         float64         `json:"index"`
	Name          string          `json:"name"`
	Skills        []Skills        `json:"skills"`
	URL           string          `json:"url"`
}
type Skills struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
