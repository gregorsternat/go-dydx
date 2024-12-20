package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func NewListModel(items []list.Item, title string, createNextModel newNextModel) tea.Model {
	log.Printf("title1: %s", title)
	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = title
	log.Printf("title2: %s", l.Title)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return listModel{
		list:            l,
		createNextModel: createNextModel,
		help:            help.New(),
		quitting:        false,
	}
}

func (m listModel) Init() tea.Cmd {
	return nil
}
