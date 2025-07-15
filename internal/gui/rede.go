package gui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"

    c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type NetworkScreen struct {
    Window fyne.Window
}

func NewNetworkScreen(w fyne.Window) *NetworkScreen {
    return &NetworkScreen{Window: w}
}

func (s *NetworkScreen) Show() {
    btnEven := widget.NewButton("Rede Par", func() {
        NewEvenScreen(s.Window).Show()
    })
    btnOdd := widget.NewButton("Rede Ímpar", func() {
        NewOddScreen(s.Window).Show()
    })

    center := c.BuildButtonGroup(
        btnEven,
        btnOdd,
    )

    footer := c.BuildFooter(
        func() { NewInstallationScreen(s.Window).Show() },
        func() { NewHomeScreen(s.Window).Show() },
    )

    c.SetScreenContent(
        s.Window,
        "Estrutura de Rede",
        center,
        footer,
    )
}
