package fake_progress

import (
	"github.com/charmbracelet/bubbles/progress"
	"time"
)

type FakeProgressModel struct {
	percent  float64
	progress progress.Model
}

type tickMsg time.Time
