$destino = Join-Path ([Environment]::GetFolderPath("Desktop")) "TempInstallers"
if (-Not (Test-Path $destino)) {
    New-Item -Path $destino -ItemType Directory | Out-Null
}

$programas = @{
    "Google_Chrome"    = "https://dl.google.com/chrome/install/latest/chrome_installer.exe"
    "Firefox"          = "https://download.mozilla.org/?product=firefox-latest&os=win64&lang=pt-BR"
    "7Zip"             = "https://www.7-zip.org/a/7z2301-x64.exe"
    "Google_Drive"     = "https://dl.google.com/drive-file-stream/GoogleDriveSetup.exe"
    "TeamViewer_Host"  = "https://download.teamviewer.com/download/TeamViewer_Host_Setup_x64.exe"
    "Google_Earth"     = "https://dl.google.com/dl/earth/client/advanced/current/googleearthprowin-7.3.6-x64.exe"
    "Driver_Booster"   = "https://cdn.iobit.com/dl/driver_booster_setup.exe"
    "Adobe_Reader"     = "https://ardownload2.adobe.com/pub/adobe/acrobat/win/AcrobatDC/2500120531/AcroRdrDCx642500120531_MUI.exe"
}

foreach ($nome in $programas.Keys) {
    $url = $programas[$nome]
    $caminho = Join-Path $destino "$nome.exe"

    Write-Host "`nbaixando $nome..."
    
    try {
        Start-BitsTransfer -Source $url -Destination $caminho
        Write-Host "$nome foi baixado com sucesso."
    } catch {
        Write-Warning "falha ao baixar ${nome}. Erro: $($_.Exception.Message)"
        continue
    }

    if (-Not (Test-Path $caminho)) {
        Write-Warning "${nome} não foi baixado corretamente. arquivo não encontrado."
        continue
    }

    Write-Host "iniciando instalação de ${nome}..."

    switch ($nome) {
        "Driver_Booster" {
            Start-Process -FilePath $caminho
            continue
        }
        "Google_Drive" {
            Start-Process -FilePath $caminho -ArgumentList "--silent", "--desktop_shortcut=true" -Wait
            continue
        }
    }

    $psi = New-Object System.Diagnostics.ProcessStartInfo
    $psi.FileName = $caminho

    switch ($nome) {
        "7Zip" {
            $psi.Arguments = "/S"
        }
        "Google_Earth" {
            $psi.Arguments = ""
        }
        "Adobe_Reader" {
            $psi.Arguments = "/sAll /rs /msi /norestart /quiet"
        }
        Default {
            $psi.Arguments = "/silent"
        }
    }

    $psi.Verb = "runas"
    $psi.UseShellExecute = $true

    try {
        $process = [System.Diagnostics.Process]::Start($psi)
        $process.WaitForExit()
        Write-Host "instalação de ${nome} concluído."
    } catch {
        Write-Warning "falha ao instalar ${nome}: $($_.Exception.Message)"
    }
}

Write-Host "`ntodos os programas foram baixados e instalados."

Start-Sleep -Seconds 5
if (Test-Path $destino) {
    try {
        Remove-Item -Path $destino -Recurse -Force
    } catch {
        Write-Warning "falha ao remover a pasta TempInstallers: $($_.Exception.Message)"
    }
}