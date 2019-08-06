package spells

type Spells struct {
	CastingTime   string       `json:"casting_time"`
	Classes       []Classes    `json:"classes"`
	Components    []string     `json:"components"`
	Concentration string       `json:"concentration"`
	Desc          []string     `json:"desc"`
	Duration      string       `json:"duration"`
	HigherLevel   []string     `json:"higher_level"`
	ID            string       `json:"_id"`
	Index         float64      `json:"index"`
	Level         float64      `json:"level"`
	Material      string       `json:"material"`
	Name          string       `json:"name"`
	Page          string       `json:"page"`
	Range         string       `json:"range"`
	Ritual        string       `json:"ritual"`
	School        School       `json:"school"`
	Subclasses    []Subclasses `json:"subclasses"`
	URL           string       `json:"url"`
}
type Classes struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type School struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Subclasses struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
