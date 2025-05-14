package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	status int
	page   int
)

const (
	todo status = iota
	inProgress
	done
)

const (
	tasks page = iota
	input
)

const (
	widthDivisor  = 3
	heightDivisor = 2
)

var (
	models      []tea.Model
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func main() {
	models = []tea.Model{newKanban(), newForm(todo)}

	p := tea.NewProgram(models[tasks])
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
