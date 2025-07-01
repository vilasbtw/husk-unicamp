package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/vilasbtw/husk-unicamp/internal/gui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Husk")

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(600, 500))

	gui.ShowHome(w)

	w.ShowAndRun()
}
