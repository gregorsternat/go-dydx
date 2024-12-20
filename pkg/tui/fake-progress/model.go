package fake_progress

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*300, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m FakeProgressModel) Init() tea.Cmd {
	return tickCmd()
}

func NewFakeProgressModel() tea.Model {
	return FakeProgressModel{
		progress: progress.New(progress.WithDefaultGradient()),
	}
}
