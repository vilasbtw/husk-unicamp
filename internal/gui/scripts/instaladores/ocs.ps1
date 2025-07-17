param(
    [string]$usuario,
    [string]$senha
)

# Alterar a $origem
$origem = "\\servidor\caminho\OCS.zip"
$destino = [Environment]::GetFolderPath("Desktop")

# Alterar o caminho
net use "\\servidor\caminho\OCS" /user:NETUNO\$usuario $senha

Copy-Item -Path $origem -Destination $destino -Force

# Alterar o caminho
net use "\\servidor\caminho\OCS" /delete