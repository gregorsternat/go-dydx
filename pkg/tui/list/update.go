package list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	text_inputs "github.com/gregorsternat/staking-validator-monitoring-tool/pkg/tui/text-inputs"
)

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			_, ok := m.list.SelectedItem().(Item)
			if ok {
				items := []list.Item{
					Item("BeraChain"),
					Item("Solana"),
				}
				//				nextModel := m.createNextModel(string(i))
				return NewListModel(items, "Which network do you want to chose ?", text_inputs.NewTextInputsModel), nil
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
