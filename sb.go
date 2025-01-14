package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	choices []string
	cursor int
	selected map[int]struct{}


}

// convert to var?
func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// keys indices of choices
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
    // initial IO - currently none
    return nil
}

func main {

}
