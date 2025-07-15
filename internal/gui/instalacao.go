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
    center := c.BuildButtonGroup(
        widget.NewButton("Configurar rede", func() {
            NewNetworkScreen(s.Window).Show()
        }),
        widget.NewButton("Instalar programas", func() {
            // NewDownloadScreen(s.Window).Show()
        }),
        widget.NewButton("Configurar TeamViewer", func() {
            // NewTeamViewerScreen(s.Window).Show()
        }),
    )

    footer := c.BuildFooter(
        func() { NewHomeScreen(s.Window).Show() },
        func() { NewHomeScreen(s.Window).Show() },
    )

    c.SetScreenContent(
        s.Window,
        "Instalação",
        center,
        footer,
    )
}