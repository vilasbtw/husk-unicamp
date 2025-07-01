package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowHome(w fyne.Window) {
	btnNetwork := widget.NewButton("Configurar Rede", func() {
		ShowNetwork(w)
	})

	btnSoftwares := widget.NewButton("Baixar Softwares", func() {
		// TO-DO
	})

	btnUpdates := widget.NewButton("Verificar Atualizações", func() {
		// TO-DO
	})

	buttonsContainer := container.NewVBox(
		btnNetwork,
		btnSoftwares,
		btnUpdates,
	)

	content := container.NewCenter(buttonsContainer)

	w.SetContent(content)
}