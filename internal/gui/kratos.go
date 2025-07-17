package gui

import (
	"os"
	"os/exec"
	"path/filepath"

	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	_ "embed"

)

//go:embed scripts/instaladores/kaspersky.ps1
var kasperskyScript []byte

type KratosScreen struct {
	Window fyne.Window
}

func NewKratosScreen(w fyne.Window) *KratosScreen {
	return &KratosScreen{Window: w}
}

func (s *KratosScreen) Show() {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Matrícula (sem o PREFEITURA\\)")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Senha")

	btnOK := widget.NewButton("Instalar", func() {
		usuario := usernameEntry.Text
		senha := passwordEntry.Text
		if usuario == "" || senha == "" {
			dialog.ShowInformation("Erro", "Preencha usuário e senha", s.Window)
			return
		}
		s.runCopyScript(usuario, senha)
	})
	btnOK.Importance = widget.HighImportance

	center := c.BuildButtonGroup(
		widget.NewLabel("Digite suas credenciais para acessar o instalador do Kaspersky:"),
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
		"Kaspersky - Kratos",
		center,
		footer,
	)

}

func (s *KratosScreen) runCopyScript(usuario, senha string) {
	temp := os.TempDir()
	scriptPath := filepath.Join(temp, "kaspersky.ps1")

	_ = os.WriteFile(scriptPath, kasperskyScript, 0o644)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath,
		"-usuario", usuario,
		"-senha", senha,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

	dialog.ShowInformation("Sucesso", "Instalador copiado para a área de trabalho.", s.Window)
	_ = output
}