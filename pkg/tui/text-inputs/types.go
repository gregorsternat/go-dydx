package text_inputs

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

type textInputsModel struct {
	focusIndex     int
	inputs         []textinput.Model
	cursorMode     cursor.Mode
	previousChoice string
	quitting       bool
}
