package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type newNextModel func(choice string) tea.Model

type listModel struct {
	list            list.Model
	createNextModel newNextModel
	help            help.Model
	quitting        bool
}

type Item string

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}
