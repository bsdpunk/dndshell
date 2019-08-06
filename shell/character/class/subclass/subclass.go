package Subclass

type Subclass struct {
	Class          Class        `json:"class"`
	Desc           []string     `json:"desc"`
	Features       []Features   `json:"features"`
	Index          float64      `json:"index"`
	Name           string       `json:"name"`
	Spells         []Spells     `json:"spells"`
	SubclassFlavor string       `json:"subclass_flavor"`
	Subclasses     []Subclasses `json:"Subclasses"`
	URL            string       `json:"url"`
}
type Class struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Features struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Prerequisites struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
type Spells struct {
	Prerequisites []Prerequisites `json:"prerequisites"`
	Spell         Spell           `json:"spell"`
}
type Subclasses struct {
	Subclasses []Subclass
}
