package logger

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/lipgloss"
)

func init() {
	styles := log.DefaultStyles()

	// Customize the Debug level style
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")). // Red background
		Foreground(lipgloss.Color("0"))    // Black text

	// Configure the default logger
	log.SetOutput(os.Stdout)
	log.SetStyles(styles)
	log.SetLevel(log.DebugLevel)
}
