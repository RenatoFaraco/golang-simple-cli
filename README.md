# Interactive Go CLI (`golang-simple-cli`)

Uma **CLI interativa desenvolvida em Go (Golang)** utilizando o framework **Cobra** e estilizada com **Fatih Color**. Projetada para executar no **WSL/Linux**, esta ferramenta oferece utilitários de rede em tempo real e um painel administrativo para o **GopherCRM**.

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)](https://golang.org/)
[![CLI Framework](https://img.shields.io/badge/Framework-Cobra-14A073?style=flat)](https://github.com/spf13/cobra)
[![Repository](https://img.shields.io/badge/GitHub-RenatoFaraco%2Fgolang--simple--cli-181717?style=flat&logo=github)](https://github.com/RenatoFaraco/golang-simple-cli)

---

## Comandos e Funcionalidades

| Comando | Tipo | Descrição e Recursos |
| :--- | :--- | :--- |
| `ip` | **Rede** | Exibe o IP Local (WSL) e o IP Público de internet com checagem de status online/offline |
| `wifi` | **Rede** | Escaneia SSIDs e sinal Wi-Fi próximo com indicativos visuais de qualidade (🟢/🟡/🔴) |
| `clear` | **Sistema** | Executa a limpeza nativa da tela do terminal (`cls`/`clear`) |
| `crm login` | **GopherCRM** | Autentica o administrador no GopherCRM e gera o token de sessão JWT |
| `exit` / `quit` | **Sessão** | Finaliza a execução da CLI interativa |

---

## Instalação e Execução Rápida

```bash
# Clone o repositório oficial
git clone [https://github.com/RenatoFaraco/golang-simple-cli.git](https://github.com/RenatoFaraco/golang-simple-cli.git)
cd golang-simple-cli

# Instale as dependências
go mod tidy

# Inicie o shell interativo
go run .