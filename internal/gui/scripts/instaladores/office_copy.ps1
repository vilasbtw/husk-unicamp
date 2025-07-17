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

# Alterar conforme sua necessidade.
$basePath = "\\servidor\caminho"

switch ($versao) {
    "2010" {
        $origem = Join-Path $basePath "2010_instalador"
    }
    "2013" {
        $origem = Join-Path $basePath "Office_Professional_Plus_2013_x86_x64_pt-br"
    }
    "2016" {
        $origem = Join-Path $basePath "OFFICE_2016_64"
    }
    "2019" {
        $origem = Join-Path $basePath "2019 - 64bits"
    }
    "2021" {
        $origem = Join-Path $basePath "Office Professional Plus 2021 - 64bits"
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
} catch {
    Write-Error "erro ao copiar a pasta: $_"
}

net use "$basePath" /delete