package text_inputs

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func NewTextInputsModel(choice string) tea.Model {
	m := textInputsModel{
		inputs: make([]textinput.Model, 1),
	}
	m.previousChoice = choice

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 64

		switch i {
		case 0:
			t.Placeholder = "Address"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.inputs[i] = t
	}

	return m
}

func (m textInputsModel) Init() tea.Cmd {
	return textinput.Blink
}
