package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	status      status
	title       textinput.Model
	description textarea.Model
}

func newForm(state status) *Form {
	form := &Form{status: state, description: textarea.New()}
	form.title = textinput.New()
	form.title.Focus()
	return form
}

func (m Form) Init() tea.Cmd {
	return textinput.Blink
}

func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, QuitKeys) {
			return m, tea.Quit
		}
		switch msg.String() {
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.description.Focus()
				return m, textarea.Blink
			} else {
				// switch to previous model, add task
				models[input] = m
				return models[tasks], m.NewTask
			}
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else {
		m.description, cmd = m.description.Update(msg)
		return m, cmd
	}
}

func (m Form) NewTask() tea.Msg {
	task := Task{status: m.status, title: m.title.Value(), description: m.description.Value()}
	log.Print(task)
	return task
}

func (m Form) helpMenu() string {
	var msg string
	if m.title.Focused() {
		msg = "next"
	} else {
		msg = "submit"
	}
	return helpStyle.Render(fmt.Sprintf("enter: %s", msg))
}

func (m Form) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.description.View(), m.helpMenu())
}
