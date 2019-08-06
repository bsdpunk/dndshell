package feats

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
