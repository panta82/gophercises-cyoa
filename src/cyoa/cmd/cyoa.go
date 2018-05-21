package main

import (
	"fmt"
	"cyoa"
)

func main() {
	story := cyoa.LoadStory()
	fmt.Println(story)
}