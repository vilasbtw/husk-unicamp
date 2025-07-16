package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
	"github.com/vilasbtw/husk-unicamp/internal/state"
)

type NetunoLoginScreen struct {
	Window fyne.Window
}

func NewNetunoLoginScreen(w fyne.Window) *NetunoLoginScreen {
	return &NetunoLoginScreen{Window: w}
}

func (s *NetunoLoginScreen) Show() {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Usuário (sem o NETUNO\\)")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Senha")

	btnOK := widget.NewButton("OK", func() {
		usuario := usernameEntry.Text
		senha := passwordEntry.Text

		state.UsuarioNetuno = usuario
		state.SenhaNetuno = senha

		NewNetunoScreen(s.Window).Show()
	})
	btnOK.Importance = widget.HighImportance

	center := c.BuildButtonGroup(
		widget.NewLabel("Digite suas credenciais para acessar os aplicativos da Netuno:"),
		usernameEntry,
		passwordEntry,
		btnOK,
	)

	footer := c.BuildFooter(
		func() { NewDownloadScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

	c.SetScreenContent(
		s.Window,
		"Login - Netuno",
		center,
		footer,
	)
}