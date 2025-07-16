package gui

import (
    _ "embed"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

//go:embed scripts/par.ps1
var evenScript []byte

//go:embed scripts/impar.ps1
var oddScript []byte

func runNetworkSetup(script []byte, filename, host string) {
    targetDir := "C:\\TempInstallers"

    if err := os.MkdirAll(targetDir, 0o755); err != nil {
        return
    }

    target := filepath.Join(targetDir, filename)
    content := strings.ReplaceAll(string(script), "HOST_PLACEHOLDER", host)
    
    if err := os.WriteFile(target, []byte(content), 0o644); err != nil {
        return
    }

    _ = exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", target).Run()
}

func runEvenNetworkSetup(host string) {
    runNetworkSetup(evenScript, "par.ps1", host)
}

func runOddNetworkSetup(host string) {
    runNetworkSetup(oddScript, "impar.ps1", host)
}