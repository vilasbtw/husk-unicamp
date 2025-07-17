param(
    [string]$usuario,
    [string]$senha
)

# Alterar a $origem
$origem = "\\servidor\caminho\installer.exe"
$destino = [Environment]::GetFolderPath("Desktop")

# Alterar o caminho
net use "\\servidor\caminho" /user:PREFEITURA\$usuario $senha

Copy-Item -Path $origem -Destination $destino -Force

# Alterar o caminho
net use "\\servidor\caminho" /delete