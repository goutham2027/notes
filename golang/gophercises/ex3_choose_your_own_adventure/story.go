package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title          string          `json:"title"`
	Paragraphs     []string        `json:"story"`
	ArcSuggestions []ArcSuggestion `json:"options"`
}

type ArcSuggestion struct {
	SuggestionText string `json:"text"`
	ArcName        string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var arcs Story

	if err := d.Decode(&arcs); err != nil {
		return nil, err
	}
	return arcs, nil
}
