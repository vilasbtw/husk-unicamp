package gui

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/vilasbtw/husk-unicamp/internal/utils"

	_ "embed"
)

//go:embed scripts/par.ps1
var evenScript []byte

//go:embed scripts/impar.ps1
var oddScript []byte

func runNetworkSetup(script []byte, filename, host string) {
	targetDir := "C:\\TempInstallers"

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		utils.LogToFile("Erro ao criar pasta TempInstallers: " + err.Error())
		return
	}

	target := filepath.Join(targetDir, filename)
	content := strings.ReplaceAll(string(script), "HOST_PLACEHOLDER", host)

	if err := os.WriteFile(target, []byte(content), 0o644); err != nil {
		utils.LogToFile("Erro ao salvar "+filename+": " + err.Error())
		return
	}

	utils.LogToFile("Executando script de rede: " + filename + " com host " + host)

	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", target)
	output, err := cmd.CombinedOutput()
	utils.LogToFile("Saída do script "+filename+":\n" + string(output))
	if err != nil {
		utils.LogToFile("Erro ao executar "+filename+": " + err.Error())
	}
}

func runEvenNetworkSetup(host string) {
	runNetworkSetup(evenScript, "par.ps1", host)
}

func runOddNetworkSetup(host string) {
	runNetworkSetup(oddScript, "impar.ps1", host)
}