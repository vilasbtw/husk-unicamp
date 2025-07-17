package gui

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vilasbtw/husk-unicamp/internal/utils"
	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	_ "embed"
)

//go:embed scripts/instaladores/installer.ps1
var installerScript []byte

type DownloadScreen struct {
	Window fyne.Window
}

func NewDownloadScreen(w fyne.Window) *DownloadScreen {
	return &DownloadScreen{Window: w}
}

func (s *DownloadScreen) Show() {
	center := c.BuildButtonGroup(
		widget.NewButton("Padrão", func() {
			s.runInstallerScript()
		}),
		widget.NewButton("Netuno", func() {
			NewNetunoLoginScreen(s.Window).Show()
		}),
		widget.NewButton("Kratos", func() {
			NewKratosScreen(s.Window).Show()
		}),
	)

	footer := c.BuildFooter(
		func() { NewHomeScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

	c.SetScreenContent(
		s.Window,
		"Instalação",
		center,
		footer,
	)
}

func (s *DownloadScreen) runInstallerScript() {
	targetDir := filepath.Join(os.Getenv("USERPROFILE"), "Desktop", "TempInstallers")
	target := filepath.Join(targetDir, "installer.ps1")

	err := os.MkdirAll(targetDir, 0o755)
	if err != nil {
		utils.LogToFile("Erro ao criar pasta TempInstallers: " + err.Error())
		return
	}

	err = os.WriteFile(target, installerScript, 0o644)
	if err != nil {
		utils.LogToFile("Erro ao salvar installer.ps1: " + err.Error())
		return
	}

	utils.LogToFile("Executando installer.ps1 em: " + target)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", target)
	output, err := cmd.CombinedOutput()
	utils.LogToFile("Saída do comando installer.ps1:\n" + string(output))
	if err != nil {
		utils.LogToFile("Erro ao executar installer.ps1: " + err.Error())
	}
}