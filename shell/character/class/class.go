package class

type Class struct {
	ClassLevels        ClassLevels          `json:"class_levels"`
	HitDie             float64              `json:"hit_die"`
	ID                 string               `json:"_id"`
	Index              float64              `json:"index"`
	Name               string               `json:"name"`
	Proficiencies      []Proficiencies      `json:"proficiencies"`
	ProficiencyChoices []ProficiencyChoices `json:"proficiency_choices"`
	SavingThrows       []SavingThrows       `json:"saving_throws"`
	StartingEquipment  StartingEquipment    `json:"starting_equipment"`
	Subclasses         []Subclasses         `json:"subclasses"`
	URL                string               `json:"url"`
}
type ClassLevels struct {
	Class string `json:"class"`
	URL   string `json:"url"`
}
type From struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Proficiencies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type ProficiencyChoices struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type SavingThrows struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingEquipment struct {
	Class string `json:"class"`
	URL   string `json:"url"`
}
type Subclasses struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Classes struct {
	Classes []Class `json:"classes"`
}
