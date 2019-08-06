package race

type Race struct {
	AbilityBonusOptions        AbilityBonusOptions        `json:"ability_bonus_options"`
	AbilityBonuses             []AbilityBonuses           `json:"ability_bonuses"`
	Age                        string                     `json:"age"`
	Alignment                  string                     `json:"alignment"`
	Index                      float64                    `json:"index"`
	LanguageDesc               string                     `json:"language_desc"`
	LanguageOptions            LanguageOptions            `json:"language_options"`
	Languages                  []Languages                `json:"languages"`
	Name                       string                     `json:"name"`
	Races                      []Races                    `json:"Races"`
	Size                       string                     `json:"size"`
	SizeDescription            string                     `json:"size_description"`
	Speed                      float64                    `json:"speed"`
	StartingProficiencies      []StartingProficiencies    `json:"starting_proficiencies"`
	StartingProficiencyOptions StartingProficiencyOptions `json:"starting_proficiency_options"`
	Subraces                   []Subraces                 `json:"subraces"`
	TraitOptions               TraitOptions               `json:"trait_options"`
	Traits                     []Traits                   `json:"traits"`
	URL                        string                     `json:"url"`
}
type AbilityBonusOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type AbilityBonuses struct {
	Bonus float64 `json:"bonus"`
	Name  string  `json:"name"`
	URL   string  `json:"url"`
}
type From struct {
	Bonus float64 `json:"bonus"`
	Name  string  `json:"name"`
	URL   string  `json:"url"`
}
type LanguageOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Languages struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingProficiencies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type StartingProficiencyOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Subraces struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type TraitOptions struct {
	Choose float64 `json:"choose"`
	From   []From  `json:"from"`
	Type   string  `json:"type"`
}
type Traits struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Races struct {
	races []Race
}
