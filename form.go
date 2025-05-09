package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	status      status
	title       textinput.Model
	description textarea.Model
	edited      bool
}

func newForm(state status) *Form {
	form := &Form{status: state, title: textinput.New(), description: textarea.New(), edited: false}
	form.title.Focus()
	return form
}

func newEditForm(task Task) *Form {
	form := &Form{status: task.status, title: textinput.New(), description: textarea.New(), edited: true}
	form.title.SetValue(task.title)
	form.description.SetValue(task.description)
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
		switch msg.String() {
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.description.Focus()
				return m, textarea.Blink
			} else {
				// switch to previous model, add task
				models[input] = m
				if m.edited {
					return models[tasks], m.EditedTask
				}
				return models[tasks], m.NewTask
			}
		case "esc":
			return models[tasks], nil
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
	task := Task{status: m.status, title: m.title.Value(), description: m.description.Value(), edited: false}
	return task
}

func (m Form) EditedTask() tea.Msg {
	task := Task{status: m.status, title: m.title.Value(), description: m.description.Value(), edited: true}
	return task
}

func (m Form) helpMenu() string {
	var msg string
	if m.title.Focused() {
		msg = "next"
	} else {
		msg = "submit"
	}
	return helpStyle.Render(fmt.Sprintf("enter: %s • esc cancel", msg))
}

func (m Form) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.description.View(), m.helpMenu())
}
