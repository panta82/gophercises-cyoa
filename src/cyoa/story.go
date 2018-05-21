package cyoa

import (
	"io/ioutil"
	"encoding/json"
)

const storyFileName = "gopher.json"
const defaultStartingArc = "intro"

type StoryOption struct {
	Text string
	Arc string
}

type StoryArc struct {
	Name string
	Title string
	Paragraphs []string `json:"story"`
	Options []StoryOption
}

type Story struct {
	StartingArc string
	Arcs map[string]StoryArc
}

func LoadStory() Story {
	fileName := getAssetPath(storyFileName)
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Failed to load story from " + fileName + ": " + err.Error())
	}

	var dict map[string]StoryArc
	err = json.Unmarshal(raw, &dict)
	if err != nil {
		panic("Failed to parse story file " + fileName + ": " + err.Error())
	}

	var story Story
	story.Arcs = make(map[string]StoryArc)
	for name, arc := range dict {
		arc.Name = name
		story.Arcs[name] = arc
	}
	story.StartingArc = defaultStartingArc

	return story
}