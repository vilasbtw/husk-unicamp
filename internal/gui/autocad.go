package gui

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vilasbtw/husk-unicamp/internal/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_ "embed"
)

//go:embed scripts/instaladores/autocad_copy.ps1
var autocadCopyScript []byte

type AutocadSelectionScreen struct {
	Window fyne.Window
	Versao string
}

func NewAutocadSelectionScreen(w fyne.Window) *AutocadSelectionScreen {
	return &AutocadSelectionScreen{Window: w}
}

func (s *AutocadSelectionScreen) Show() {
	versions := []string{"2016", "2017", "2022", "2024", "2025"}
	btns := []fyne.CanvasObject{}

	for _, v := range versions {
		version := v
		btn := widget.NewButton(version, func() {
			s.Versao = version
		})
		btns = append(btns, btn)
	}

	installBtn := widget.NewButton("Instalar", func() {
		if s.Versao == "" {
			dialog.ShowInformation("Erro", "Nenhuma versão selecionada.", s.Window)
			return
		}
		s.runCopyScript()
	})
	installBtn.Importance = widget.HighImportance

	btnBack := widget.NewButton("Voltar", func() {
		NewNetunoScreen(s.Window).Show()
	})

	btnHome := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		NewHomeScreen(s.Window).Show()
	})

	footer := container.NewHBox(layout.NewSpacer(), btnBack, btnHome)

	content := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(widget.NewLabel("Selecione a versão do AutoCAD para instalar:")),
		container.NewCenter(container.NewHBox(btns...)),
		installBtn,
		layout.NewSpacer(),
	)

	s.Window.SetContent(container.NewBorder(nil, footer, nil, nil, content))
}

func (s *AutocadSelectionScreen) runCopyScript() {
	temp := os.TempDir()
	scriptPath := filepath.Join(temp, "autocad_copy.ps1")

	err := os.WriteFile(scriptPath, autocadCopyScript, 0644)
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

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

	dialog.ShowInformation("Sucesso", "Versão do AutoCAD copiada com sucesso para a área de trabalho.", s.Window)
	_ = output
}
