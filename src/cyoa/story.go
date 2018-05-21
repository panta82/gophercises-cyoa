package cyoa

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

const storyFileName = "gopher.json"
const startingArc = "intro"

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
	CurrentArc string
	Arcs map[string]StoryArc
}

func getStoryFileName() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Failed to determine directory of the executable")
	}

	return filepath.Join(dir, storyFileName)
}

func LoadStory() Story {
	raw, err := ioutil.ReadFile(getStoryFileName())
	if err != nil {
		panic("Failed to load story from " + getStoryFileName() + ": " + err.Error())
	}

	var dict map[string]StoryArc
	err = json.Unmarshal(raw, &dict)
	if err != nil {
		panic("Failed to parse story file " + getStoryFileName() + ": " + err.Error())
	}

	var story Story
	story.Arcs = make(map[string]StoryArc)
	for name, arc := range dict {
		arc.Name = name
		story.Arcs[name] = arc
	}
	story.CurrentArc = startingArc

	return story
}