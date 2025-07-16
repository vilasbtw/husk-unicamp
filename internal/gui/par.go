package gui

import (
	"fyne.io/fyne/v2"
	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
)

type EvenScreen struct {
	Window   fyne.Window
	Selected string
}

func NewEvenScreen(w fyne.Window) *EvenScreen {
	return &EvenScreen{
		Window: w, 
		// Altere o host padrão para a rede par.
		Selected: "01",
	}
}

func (s *EvenScreen) Show() {
	c.ShowNetworkConfig(
		s.Window,
		"Rede Par",
		// Altere aqui para definir o prefixo do IP da rede par
		"102.102.102.",
		// Botões para definir os últimos octetos (hosts)
		[]string{"01", "02", "03", "04"},
		&s.Selected,
		runEvenNetworkSetup,
		func() { NewNetworkScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

}