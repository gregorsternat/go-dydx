package main

import (
	"context"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	tuilist "github.com/gregorsternat/staking-validator-monitoring-tool/pkg/tui/list"
	readme "github.com/gregorsternat/staking-validator-monitoring-tool/pkg/tui/read-me"
	text_inputs "github.com/gregorsternat/staking-validator-monitoring-tool/pkg/tui/text-inputs"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func startOrReadMe(choice string) tea.Model {
	if choice == "Start" {
		items := []list.Item{
			tuilist.Item("BeraChain"),
			tuilist.Item("Solana"),
		}

		return tuilist.NewListModel(items, "Which network do you want to chose ?", text_inputs.NewTextInputsModel)
	}
	m, err := readme.NewReadMeModel()
	if err != nil {
		return nil
	}
	return m
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	items := []list.Item{
		tuilist.Item("Start"),
		tuilist.Item("Read me"),
	}

	f, _ := tea.LogToFile("log.txt", "debug")
	defer f.Close()

	log.Printf("Starting the program...")
	p := tea.NewProgram(tuilist.NewListModel(
		items, "What do you want to do?", startOrReadMe), tea.WithContext(ctx),
	)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
