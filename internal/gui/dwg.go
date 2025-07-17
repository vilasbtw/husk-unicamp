package gui

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vilasbtw/husk-unicamp/internal/state"
	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"

	_ "embed"
)

//go:embed scripts/instaladores/dwg_copy.ps1
var dwgCopyScript []byte

type DwgSelectionScreen struct {
	Window fyne.Window
	Versao string
}

func NewDwgSelectionScreen(w fyne.Window) *DwgSelectionScreen {
	return &DwgSelectionScreen{Window: w}
}

func (s *DwgSelectionScreen) Show() {
	versions := []string{"2016", "2018", "2022", "2023"}
	btns := make([]fyne.CanvasObject, len(versions))

	for i, v := range versions {
		version := v
		btns[i] = widget.NewButton(version, func() {
			s.Versao = version
		})
	}

	installBtn := widget.NewButton("Instalar", func() {
		if s.Versao == "" {
			dialog.ShowInformation("Erro", "Nenhuma versão selecionada.", s.Window)
			return
		}
		s.runCopyScript()
	})
	installBtn.Importance = widget.HighImportance

	center := c.BuildButtonGroup(
		widget.NewLabel("Selecione a versão do DWG TrueView:"),
		container.NewCenter(container.NewHBox(btns...)),
		installBtn,
	)

	footer := c.BuildFooter(
		func() { NewNetunoScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

	c.SetScreenContent(
		s.Window,
		"Instalação do DWG Viewer",
		center,
		footer,
	)
}

func (s *DwgSelectionScreen) runCopyScript() {
	temp := os.TempDir()
	scriptPath := filepath.Join(temp, "dwg_copy.ps1")

	_ = os.WriteFile(scriptPath, dwgCopyScript, 0o644)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath,
		"-usuario", state.UsuarioNetuno,
		"-senha", state.SenhaNetuno,
		"-versao", s.Versao,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

	dialog.ShowInformation("Sucesso", "Versão do DWG Viewer copiada com sucesso para a área de trabalho.", s.Window)
	_ = output
}