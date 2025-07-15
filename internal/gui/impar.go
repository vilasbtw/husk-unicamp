package gui

import (
    "fyne.io/fyne/v2"
    c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type OddScreen struct {
    Window   fyne.Window
    Selected string
}

func NewOddScreen(w fyne.Window) *OddScreen {
    return &OddScreen{
        Window:   w,
        Selected: "41",
    }
}

func (s *OddScreen) Show() {
    c.ShowNetworkConfig(
        s.Window,
        "Rede Ímpar",
        "101.101.101.",
        []string{"05", "06"},
        &s.Selected,
        // TODO runOddNetworkSetup,                          
        func() { NewNetworkScreen(s.Window).Show() },
        func() { NewHomeScreen(s.Window).Show() },
    )
}
