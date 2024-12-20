package list

import "log"

func (m listModel) View() string {
	if m.quitting {
		return quitTextStyle.Render("Quitting...")
	}
	log.Printf("title: %s", m.list.Title)

	return "\n" + m.list.View()
}
