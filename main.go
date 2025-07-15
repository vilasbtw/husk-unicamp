package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
    "github.com/vilasbtw/husk-unicamp/internal/gui"
)

func main() {
	myApp := app.New()

	
	w := myApp.NewWindow("Husk")

	home := gui.NewHomeScreen(w)
	
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(450, 400))

	home.Show()
	w.ShowAndRun()
}