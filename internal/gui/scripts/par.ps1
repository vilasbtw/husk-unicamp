# verifica interface de rede: "Wifi", "Ethernet" ou "Ethernet 2"
$interface = Get-NetAdapter |
    Where-Object { $_.Status -eq "Up" -and $_.HardwareInterface -eq $true } |
    Sort-Object -Property InterfaceMetric |
    Select-Object -First 1 -ExpandProperty Name

if (-not $interface) {
    Write-Warning "Não foi possível determinar a interface de rede."
    exit
}

Write-Host "Interface: $interface"

# altere essas informações conforme necessário.
$ip        = "102.102.102.HOST_PLACEHOLDER"
$prefix    = 2
$gateway   = "102.102.102.102"
$dns1      = "102.102.102.102"
$dns2      = "102.102.102.102"
$sufixos   = @("google.com.br", "google.com.br")

# remove ip antigo
Get-NetIPAddress -InterfaceAlias $interface -AddressFamily IPv4 -ErrorAction SilentlyContinue |
    Remove-NetIPAddress -Confirm:$false -ErrorAction SilentlyContinue

# remove gateway antigo
Get-NetRoute -InterfaceAlias $interface -DestinationPrefix "0.0.0.0/0" -ErrorAction SilentlyContinue |
    Remove-NetRoute -Confirm:$false -ErrorAction SilentlyContinue

# define o novo ip e gateway
New-NetIPAddress -InterfaceAlias $interface -IPAddress $ip -PrefixLength $prefix -DefaultGateway $gateway -AddressFamily IPv4

# define a dns
Set-DnsClientServerAddress -InterfaceAlias $interface -ServerAddresses $dns1, $dns2

# define os sufixos
Set-DnsClientGlobalSetting -SuffixSearchList $sufixos

Write-Host "`nrede par configurada com sucesso."