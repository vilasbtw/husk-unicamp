package components

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func ShowNetworkConfig(
    w fyne.Window,
    title, prefix string,
    hosts []string,
    selected *string,
    onOK func(string),
    onBack, onHome func(),
) {
    ipEntry := widget.NewEntry()
    ipEntry.SetText(prefix + *selected)
    ipEntry.Disable()

    hostContainer := container.NewHBox()
    for _, h := range hosts {
        host := h
        btn := widget.NewButton(host, func() {
            *selected = host
            ipEntry.SetText(prefix + host)
        })
        hostContainer.Add(btn)
    }

    btnOK := widget.NewButton("OK", func() {
        onOK(*selected)
    })
    btnOK.Importance = widget.HighImportance

    center := container.NewVBox(
        layout.NewSpacer(),
        ipEntry,
        container.NewCenter(hostContainer),
        container.NewCenter(btnOK),
        layout.NewSpacer(),
    )

    footer := BuildFooter(onBack, onHome)
    SetScreenContent(w, title, center, footer)
}