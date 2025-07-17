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
        $origem = Join-Path $basePath "Autocad_LT_2016"
    }
    "2017" {
        $origem = Join-Path $basePath "Autocad_LT_2017"
    }
    "2022" {
        $origem = Join-Path $basePath "AutoCAD LT 2022 - (EN)"
    }
    "2024" {
        $origem = Join-Path $basePath "AutoCAD LT 2024 - (EN)"
    }
    "2025" {
        $origem = Join-Path $basePath "Autocad_LT_2025-EN"
    }
    default {
        Write-Error "Versão inválida: $versao"
        exit 1
    }
}

$destino = [Environment]::GetFolderPath("Desktop")

net use "$basePath" /user:$usuarioCompleto $senha

try {
    Copy-Item -Path $origem -Destination $destino -Recurse -Force
    Write-Host "Cópia do AutoCAD $versao concluída com sucesso."
} catch {
    Write-Error "Erro ao copiar a pasta: $_"
}

net use "$basePath" /delete