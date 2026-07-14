package cmd

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Mostra os IPs local e público da máquina",
	Long:  `Verifica as interfaces de rede locais do WSL e faz uma requisição rápida para descobrir o IP público de internet.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌐 Verificando conexões de rede...")
		fmt.Println("--------------------------------------------------")

		localIP := getLocalIP()
		fmt.Printf("💻 IP Local (WSL):  %s\n", localIP)

		publicIP := getPublicIP()
		fmt.Printf("🌍 IP Público:       %s\n", publicIP)
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
		return "OFFLINE ❌ (Erro ao conectar à internet)"
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Erro ao ler resposta"
	}

	return string(ipBytes) + " ✅ (Online)"
}

func init() {
	RootCmd.AddCommand(ipCmd)
}