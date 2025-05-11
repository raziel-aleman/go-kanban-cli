package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Kanban struct {
	focus    status
	loaded   bool
	lists    []list.Model
	quitting bool
}

func (k *Kanban) Next() {
	if k.focus == done {
		k.focus = todo
	} else {
		k.focus++
	}
}

func (k *Kanban) Prev() {
	if k.focus == todo {
		k.focus = done
	} else {
		k.focus--
	}
}

func newKanban() Kanban {
	k := Kanban{focus: todo, loaded: false}
	return k
}

func (k *Kanban) initLists(width, height int) {
	// init list model
	defaultList := list.New([]list.Item{}, NewTaskDelegate(), width/widthDivisor-6, height/heightDivisor-2) //width/widthDivisor-6, height/heightDivisor*2)
	defaultList.SetShowHelp(true)
	defaultList.SetFilteringEnabled(false)
	defaultList.FilterInput.SetValue("")
	k.lists = []list.Model{defaultList, defaultList, defaultList}

	//add items from storage
	columns := ReadFromStorage()

	// Init To Dos
	k.lists[todo].Title = "To Do"
	k.lists[inProgress].Title = "In Progress"
	k.lists[done].Title = "Done"

	todoItems := []list.Item{}
	inProgressItems := []list.Item{}
	doneItems := []list.Item{}

	for _, value := range columns.Todo {
		task := Task{todo, value.Title, value.Description, false}
		todoItems = append(todoItems, task)
	}

	for _, value := range columns.InProgress {
		task := Task{inProgress, value.Title, value.Description, false}
		inProgressItems = append(inProgressItems, task)
	}

	for _, value := range columns.Done {
		task := Task{done, value.Title, value.Description, false}
		doneItems = append(doneItems, task)
	}

	k.lists[todo].SetItems(todoItems)
	k.lists[inProgress].SetItems(inProgressItems)
	k.lists[done].SetItems(doneItems)
}

func (k Kanban) Init() tea.Cmd {
	return nil
}

func (k Kanban) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !k.loaded {
			columnStyle.Width(msg.Width/widthDivisor - 2)
			focusedStyle.Width(msg.Width/widthDivisor - 2)
			columnStyle.Height(msg.Height / heightDivisor)
			focusedStyle.Height(msg.Height / heightDivisor)
			k.initLists(msg.Width, msg.Height)
			k.loaded = true
		}
	case tea.KeyMsg:
		if key.Matches(msg, QuitKeys) && (k.lists[k.focus].FilterState() != list.Filtering) {
			k.quitting = true
			WriteToStorage(k)
			return k, tea.Quit
		}
		switch msg.String() {
		case "f":
			if k.lists[k.focus].FilterState() != list.Filtering {
				k.lists[k.focus].SetFilteringEnabled(true)
				currList, cmd := k.lists[k.focus].Update(msg)
				k.lists[k.focus] = currList
				return k, cmd
			} else {
				return k, nil
			}
		case "right":
			k.Next()
		case "left":
			k.Prev()
		case "enter":
			k.MoveToNext()
			k.Next()
			return k, nil
		case "n":
			if k.lists[k.focus].FilterState() != list.Filtering {
				// save state of current model before switching models
				models[tasks] = k
				models[input] = newForm(k.focus)
				return models[input].Update(nil)
			}
		case "x":
			if k.lists[k.focus].FilterState() != list.Filtering {
				index := k.lists[k.focus].Index()
				k.lists[k.focus].RemoveItem(index)
			}
		case "ctrl+n":
			currentList := k.lists[k.focus]
			selectedItem := currentList.SelectedItem()

			if selectedItem != nil {
				taskToEdit, ok := selectedItem.(Task)
				if ok {
					// Save the current model before switcing models
					models[tasks] = k
					models[input] = newEditForm(taskToEdit)
					return models[input].Update(nil)
				}
			}
			// If no item selected, or cast failed, stay on the current model
			return k, nil
		}
	case Task:
		task := msg
		if task.edited {
			index := k.lists[k.focus].Index()
			return k, k.lists[task.status].SetItem(index, task)
		}
		return k, k.lists[task.status].InsertItem(len(k.lists[task.status].Items()), task)
	}
	currList, cmd := k.lists[k.focus].Update(msg)
	k.lists[k.focus] = currList
	return k, cmd
}

func (k *Kanban) MoveToNext() tea.Msg {
	selectedItem := k.lists[k.focus].SelectedItem()

	if selectedItem != nil {
		selectedTask := selectedItem.(Task)
		k.lists[selectedTask.status].RemoveItem(k.lists[k.focus].Index())
		selectedTask.Next()
		k.lists[selectedTask.status].InsertItem(len(k.lists[selectedTask.status].Items()), list.Item(selectedTask))
		k.lists[selectedTask.status].Select(len(k.lists[selectedTask.status].Items()) - 1)
	}

	return nil
}

func (k Kanban) View() string {
	if k.quitting {
		return ""
	}

	if k.loaded {
		todoView := k.lists[todo].View()
		inProgView := k.lists[inProgress].View()
		doneView := k.lists[done].View()
		switch k.focus {
		case inProgress:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				focusedStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		case done:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				columnStyle.Render(inProgView),
				focusedStyle.Render(doneView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(todoView),
				columnStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		}

	} else {
		return "Loading..."
	}
}
