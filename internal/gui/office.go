package gui

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vilasbtw/husk-unicamp/internal/state"
	"github.com/vilasbtw/husk-unicamp/internal/utils"
	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"

	_ "embed"
)

//go:embed scripts/instaladores/office_copy.ps1
var officeCopyScript []byte

type OfficeSelectionScreen struct {
	Window fyne.Window
	Versao string
}

func NewOfficeSelectionScreen(w fyne.Window) *OfficeSelectionScreen {
	return &OfficeSelectionScreen{Window: w}
}

func (s *OfficeSelectionScreen) Show() {
	versions := []string{"2010", "2013", "2016", "2019", "2021"}

	versionButtons := make([]fyne.CanvasObject, len(versions))
	for i, v := range versions {
		version := v
		versionButtons[i] = widget.NewButton(version, func() {
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
		widget.NewLabel("Selecione a versão do Microsoft Office:"),
		container.NewHBox(versionButtons...),
		installBtn,
	)

	footer := c.BuildFooter(
		func() { NewNetunoScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

	c.SetScreenContent(
		s.Window,
		"Instalação do Office",
		center,
		footer,
	)
}

func (s *OfficeSelectionScreen) runCopyScript() {
	temp := os.TempDir()
	scriptPath := filepath.Join(temp, "office_copy.ps1")

	if err := os.WriteFile(scriptPath, officeCopyScript, 0o644); err != nil {
		utils.LogToFile("Erro ao salvar office_copy.ps1: " + err.Error())
		dialog.ShowError(err, s.Window)
		return
	}

	utils.LogToFile("Executando office_copy.ps1 com versão: " + s.Versao)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath,
		"-usuario", state.UsuarioNetuno,
		"-senha", state.SenhaNetuno,
		"-versao", s.Versao,
	)

	output, err := cmd.CombinedOutput()
	utils.LogToFile("Saída do comando office_copy.ps1:\n" + string(output))
	if err != nil {
		utils.LogToFile("Erro ao executar office_copy.ps1: " + err.Error())
		dialog.ShowError(err, s.Window)
		return
	}

	dialog.ShowInformation("Sucesso", "Versão do Office copiada com sucesso para a área de trabalho.", s.Window)
}