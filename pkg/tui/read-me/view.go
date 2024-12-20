package read_me

func (m readMeModel) View() string {
	return m.viewport.View() + helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}
