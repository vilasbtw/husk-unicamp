package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowNetwork(w fyne.Window) {
	buttonEven := widget.NewButton("Rede Par", func() {
	})

	buttonOdd := widget.NewButton("Rede Ímpar", func() {
	})

	buttonBack := widget.NewButton("Voltar", func() {
		ShowHome(w)
	})

	networkButtons := container.NewHBox(
		layout.NewSpacer(),
		buttonEven,
		layout.NewSpacer(),
		buttonOdd,
		layout.NewSpacer(),
	)

	content := container.NewBorder(
		// header
		nil, 
		// footer
		container.NewHBox(layout.NewSpacer(), buttonBack), 
		// left	
		nil, 
		// right
		nil, 
		// center
		networkButtons,
	)

	w.SetContent(content)
}