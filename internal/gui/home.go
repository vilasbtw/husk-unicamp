package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type HomeScreen struct {
	Window fyne.Window
}

func NewHomeScreen(w fyne.Window) *HomeScreen {
	return &HomeScreen{Window: w}
}

func (s *HomeScreen) Show() {
	c.SetScreenContent(
		s.Window,
		"Bem-vindo ao Husk!",
		newContentButtons(s),
		nil,
	)
}

func newContentButtons(s *HomeScreen) fyne.CanvasObject {
	return c.BuildButtonGroup(
		widget.NewButton("Instalação", func() {
			NewInstallationScreen(s.Window).Show()
		}),
		widget.NewButton("Pós-instalação", func() {
			// NewPostInstallationScreen(s.Window).Show()
		}),
	)
}