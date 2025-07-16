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
        // Altere o host padrão para a rede ímpar.
        Selected: "05",
    }
}

func (s *OddScreen) Show() {
    c.ShowNetworkConfig(
        s.Window,
        "Rede Ímpar",
        // Altere aqui para definir o prefixo do IP da rede ímpar.
        "101.101.101.",
        // Botões para definir os últimos octetos (hosts)
        []string{"05", "06"},
        &s.Selected,
        runOddNetworkSetup,                          
        func() { NewNetworkScreen(s.Window).Show() },
        func() { NewHomeScreen(s.Window).Show() },
    )
}