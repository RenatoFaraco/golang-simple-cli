package cmd

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Mostra os IPs local e público da máquina",
	Long:  `Verifica as interfaces de rede locais do WSL e faz uma requisição rápida para descobrir o IP público de internet.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Criando estilos reutilizáveis para organizar o layout
		cyan := color.New(color.FgCyan).SprintFunc()
		whiteBold := color.New(color.FgWhite, color.Bold).SprintFunc()

		fmt.Println(cyan("🌐 Verificando conexões de rede..."))
		fmt.Println(cyan("--------------------------------------------------"))

		// 1. IP Local (Destacando o IP com Amarelo)
		localIP := getLocalIP()
		fmt.Printf("💻 %-18s %s\n", whiteBold("IP Local (WSL):"), color.YellowString(localIP))

		// 2. IP Público (Cores dinâmicas se estiver Online ou Offline)
		publicIP := getPublicIP()
		fmt.Printf("🌍 %-18s %s\n", whiteBold("IP Público:"), publicIP)
	},
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Não foi possível obter o IP local"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Nenhum IP ativo encontrado"
}

func getPublicIP() string {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get("https://api.ipify.org")
	if err != nil {
		// Retorna vermelho se falhar
		return color.RedString("OFFLINE ❌ (Erro ao conectar à internet)")
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return color.RedString("Erro ao ler resposta")
	}

	// Retorna verde se estiver online com o IP público
	return color.GreenString(string(ipBytes) + " ✅ (Online)")
}

func init() {
	RootCmd.AddCommand(ipCmd)
}