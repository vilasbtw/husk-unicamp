package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	w := a.NewWindow("Husk")

	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}