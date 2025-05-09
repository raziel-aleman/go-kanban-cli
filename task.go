package main

import "strings"

type Task struct {
	status      status
	title       string
	description string
	edited      bool
}

func (t Task) FilterValue() string {
	return strings.ToLower(t.title)
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t *Task) Next() {
	if t.status == done {
		t.status = todo
	} else {
		t.status++
	}
}
