package gui

import (
	"fyne.io/fyne/v2"
	"github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type EvenScreen struct {
    Window   fyne.Window
    Selected string
}

func NewEvenScreen(w fyne.Window) *EvenScreen {
    return &EvenScreen{Window: w, Selected: "31"}
}

func (s *EvenScreen) Show() {
    components.ShowNetworkConfig(
        s.Window,
        "Rede Par",
        "100.100.100.",
        []string{"01", "02", "03", "04"},
        &s.Selected,
        // TODO runEvenNetworkSetup,                      
        func() { NewNetworkScreen(s.Window).Show() },
        func() { NewHomeScreen(s.Window).Show() },
    )
}
