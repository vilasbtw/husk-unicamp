param(
    [string]$usuario,
    [string]$senha
)

$origem = "\\servidor\caminho\installer.exe"
$destino = [Environment]::GetFolderPath("Desktop")

net use "\\servidor\caminho" /user:PREFEITURA\$usuario $senha

Copy-Item -Path $origem -Destination $destino -Force

net use "\\servidor\caminho" /delete