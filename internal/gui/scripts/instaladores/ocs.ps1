param(
    [string]$usuario,
    [string]$senha
)

$origem = "\\servidor\caminho\OCS.zip"
$destino = [Environment]::GetFolderPath("Desktop")

net use "\\servidor\caminho\OCS" /user:NETUNO\$usuario $senha

Copy-Item -Path $origem -Destination $destino -Force

net use "\\servidor\caminho\OCS" /delete