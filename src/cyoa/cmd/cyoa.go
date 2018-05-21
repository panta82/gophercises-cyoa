package main

import (
	"fmt"
	"cyoa"
	"net/http"
	"log"
)

const port = "8080"

func main() {
	story := cyoa.LoadStory()
	handler := cyoa.NewStoryHandler(story)

	fmt.Println("Listening on " + port)
	err := http.ListenAndServe(":" + port, handler)
	log.Fatal(err)
}