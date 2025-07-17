package gui

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	c "github.com/vilasbtw/husk-unicamp/internal/gui/components"
	"github.com/vilasbtw/husk-unicamp/internal/state"

	_ "embed"
)

//go:embed scripts/instaladores/ocs.ps1
var ocsCopyScript []byte

type NetunoScreen struct {
	Window fyne.Window
}

func NewNetunoScreen(w fyne.Window) *NetunoScreen {
	return &NetunoScreen{Window: w}
}

func (s *NetunoScreen) Show() {
	center := c.BuildButtonGroup(
		widget.NewButton("Microsoft Office", func() {
			NewOfficeSelectionScreen(s.Window).Show()
		}),
		widget.NewButton("OCS", func() {
			s.runCopyScript(ocsCopyScript, "ocs.ps1")
		}),
		widget.NewButton("AutoCAD", func() {
			// NewAutocadSelectionScreen(s.Window).Show()
		}),
		widget.NewButton("DWG TrueView", func() {
			// NewDwgSelectionScreen(s.Window).Show()
		}),
	)

	footer := c.BuildFooter(
		func() { NewDownloadScreen(s.Window).Show() },
		func() { NewHomeScreen(s.Window).Show() },
	)

	c.SetScreenContent(
		s.Window,
		"Aplicativos da Netuno",
		center,
		footer,
	)
}

func (s *NetunoScreen) runCopyScript(script []byte, scriptName string) {
	temp := os.TempDir()
	scriptPath := filepath.Join(temp, scriptName)

	_ = os.WriteFile(scriptPath, script, 0o644)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath,
		"-usuario", state.UsuarioNetuno,
		"-senha", state.SenhaNetuno,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

	appName := strings.ToUpper(strings.TrimSuffix(scriptName, ".ps1"))
	dialog.ShowInformation("Sucesso", "Instalador do "+appName+" copiado com sucesso para a área de trabalho.", s.Window)
	_ = output
}