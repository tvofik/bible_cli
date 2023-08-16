package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func boy(passage map[int]string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(2).
		Width(200)
	fmt.Println(style.Render(passage))
}
