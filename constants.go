package main

import "github.com/charmbracelet/bubbles/key"

type (
	ErrMsg error
)

var QuitKeys = key.NewBinding(
	key.WithKeys("q", "ctrl+c"),
)
