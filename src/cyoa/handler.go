package cyoa

import (
	"net/http"
	"io/ioutil"
	"html/template"
)

const templateFileName = "story_arc.gohtml"

type StoryHandler struct {
	story Story
	arcTemplate template.Template
}

func loadTemplate() template.Template {
	fileName := getAssetPath(templateFileName)
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Failed to load html template from " + fileName + ": " + err.Error())
	}

	tmpl, err := template.New("story_arc").Parse(string(raw))
	if err != nil {
		panic("Failed to parse template at " + fileName + ": " + err.Error())
	}

	return *tmpl
}

func NewStoryHandler(story Story) StoryHandler {
	return StoryHandler{
		story: story,
		arcTemplate: loadTemplate(),
	}
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arcKey := r.URL.Path
	if arcKey == "/" {
		arcKey = h.story.StartingArc
	} else {
		arcKey = arcKey[1:]
	}

	arc, ok := h.story.Arcs[arcKey]
	if !ok {
		http.NotFound(w, r)
		return
	}

	h.arcTemplate.Execute(w, &arc)
}