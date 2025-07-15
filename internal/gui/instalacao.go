package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type InstallationScreen struct {
	Window fyne.Window
}

func NewInstallationScreen(w fyne.Window) *InstallationScreen {
	return &InstallationScreen{Window: w}
}

func (s *InstallationScreen) Show() {
	c.SetScreenContent(
		s.Window,
		"Instalação",
		newInstallationButtons(s.Window),
		newInstallationFooter(s.Window),
	)
}

func newInstallationButtons(w fyne.Window) fyne.CanvasObject {
	return c.BuildButtonGroup(
		widget.NewButton("Configurar rede", func() {
			// NewNetworkScreen(w).Show()
		}),
		widget.NewButton("Instalar programas", func() {
			// NewDownloadScreen(w).Show()
		}),
		widget.NewButton("Configurar TeamViewer", func() {
			// NewDownloadScreen(w).Show()
		}),
	)
}

func newInstallationFooter(s fyne.Window) fyne.CanvasObject {
	return c.BuildFooter(
		func() { 
			NewHomeScreen(s).Show() 
		},
		func() {
			NewHomeScreen(s).Show() 
		},
	)
}