package lang

import "fmt"

type Language struct {
	Index           int      `json:"index"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	TypicalSpeakers []string `json:"typical_speakers"`
	Script          string   `json:"script"`
	URL             string   `json:"url"`
}
type Languages struct {
	Languages []Language
}
