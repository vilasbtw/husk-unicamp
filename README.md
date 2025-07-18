# Husk

Software desenvolvido para automatizar a configuração de computadores recém-formatados da Unicamp. Ele oferece uma interface gráfica para:

- Configurar automaticamente a rede
- Instalar softwares diretamente da web ou do servidor da Unicamp
- Padronizar o ambiente com poucos cliques

## Requisitos
- [Golang](https://go.dev/doc/install)
- [Fyne](https://docs.fyne.io/started/)

## Funcionalidades

- Configuração automática de IP, máscara de rede, gateway e DNS
- Detecção da interface de rede
- Instalação de programas como Office, TeamViewer, Kaspersky e outros
- Interface gráfica intuitiva e simples
- Execução de scripts PowerShell embutidos no binário

## Gerando o executável
Se estiver usando Linux, utilize o seguinte comando para compilar o executável .exe para Windows:

```bash
GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc fyne package -os windows
```

Se estiver usando o próprio Windows, basta rodar:
```bash
fyne package -os windows
```

## Tecnologias Utilizadas

- **Go (Golang)** – linguagem principal do projeto
- **Fyne** – biblioteca gráfica para construção da UI
- **PowerShell** – scripts para configuração e instalação
