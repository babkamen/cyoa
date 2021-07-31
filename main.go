package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 1 {

	}
	jsonFilepath := "gopher.json"

	jsonData, err := ioutil.ReadFile(jsonFilepath)
	logFatal("Error while reading json file", err)
	var story map[string]StoryArc
	err = json.Unmarshal(jsonData, &story)
	logFatal("Error while converting json", err)
	if len(story) == 0 {
		log.Fatal("Empty book")
	}
	processArc("intro", story)
}

func processArc(chosenArc string, story map[string]StoryArc) {
	arc := story[chosenArc]
	fmt.Println(arc.Title)
	fmt.Printf(strings.Join(arc.Story[:], "\n"))
	fmt.Println()

	optionsCount := len(arc.Options)
	if optionsCount == 0 {
		return
	}

	choices := make(map[int]string, optionsCount)
	for i, v := range arc.Options {
		choices[i+1] = v.Arc
	}

	fmt.Println("Available options:")
	for {
		printChoices(arc)
		result := readInputFromUser()
		choice, err := strconv.Atoi(result)
		logFatal("Error while reading user input", err)
		v, b := choices[choice]
		if b {
			processArc(v, story)
			return
		}
	}

}

func printChoices(arc StoryArc) {
	for i, v := range arc.Options {
		fmt.Println(i+1, "- ", v.Text)
	}
}

func readInputFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	logFatal("Something went wrong when reading input ", err)
	result = strings.TrimSpace(result)
	return result
}

func logFatal(message string, err error) {
	if err != nil {
		log.Fatal(message, " ", err)
	}
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}
