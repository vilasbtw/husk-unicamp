package components

import (
	"image/color"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)
func VerticalSpacer(height float32) fyne.CanvasObject {
	space := canvas.NewRectangle(color.Transparent)
	space.SetMinSize(fyne.NewSize(0, height))
	return space
}

func BuildTitle(text string) fyne.CanvasObject {
	title := canvas.NewText(text, color.Black)
	title.TextStyle.Bold = true
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		VerticalSpacer(50),
		container.NewCenter(title),
	)
}

func BuildButtonGroup(buttons ...fyne.CanvasObject) fyne.CanvasObject {
	grid := container.New(layout.NewGridLayoutWithColumns(1), buttons...)

	return container.NewVBox(
		VerticalSpacer(50),
		container.NewCenter(grid),
		layout.NewSpacer(),
	)
}

func BuildFooter(onBack, onHome func()) fyne.CanvasObject {
	back := widget.NewButtonWithIcon("Voltar", theme.NavigateBackIcon(), onBack)
	home := widget.NewButtonWithIcon("", theme.HomeIcon(), onHome)

	return container.NewHBox(
		layout.NewSpacer(),
		back,
		home,
	)
}

func SetScreenContent(w fyne.Window, title string, center fyne.CanvasObject, footer fyne.CanvasObject) {
	w.SetContent(container.NewBorder(
		BuildTitle(title),
		footer,
		nil, nil,
		center,
	))
}