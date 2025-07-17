param(
    [string]$usuario,
    [string]$senha,
    [string]$versao
)

if (-not $usuario -or -not $senha -or -not $versao) {
    Write-Error "Parâmetros ausentes. Certifique-se de fornecer usuário, senha e versão."
    exit 1
}

$usuarioCompleto = "NETUNO\$usuario"
# alterar conforme sua necessidade.
$basePath = "\\servidor\caminho\AUTOCAD"

switch ($versao) {
    "2016" {
        $arquivo = "DWGTrueView_2016_Enu_64bit_v3.sfx.exe"
    }
    "2018" {
        $arquivo = "DWGTrueView_2018_ENU_64bit.sfx.exe"
    }
    "2022" {
        $arquivo = "DWGTrueView_2022_English_64bit_dlm.sfx.exe"
    }
    "2023" {
        $arquivo = "DWGTrueView_2023_English_64bit_dlm.sfx.exe"
    }
    default {
        Write-Error "Versão inválida: $versao"
        exit 1
    }
}

$origem = Join-Path $basePath $arquivo
$destino = [Environment]::GetFolderPath("Desktop")

net use "$basePath" /user:$usuarioCompleto $senha

try {
    Copy-Item -Path $origem -Destination $destino -Force
    Write-Host "Cópia do DWG Viewer $versao concluída com sucesso."
} catch {
    Write-Error "Erro ao copiar o arquivo: $_"
}

net use "$basePath" /delete