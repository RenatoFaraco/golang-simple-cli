package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "Lista as redes Wi-Fi disponíveis e a intensidade do sinal",
	Long:  `Executa o utilitário do Windows para escanear e listar os SSIDs e sinais das redes próximas.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Escaneando redes Wi-Fi próximas...")
		fmt.Println("--------------------------------------------------")

		out, err := exec.Command("cmd.exe", "/c", "netsh wlan show networks mode=bssid").Output()
		if err != nil {
			log.Fatalf("Erro ao buscar redes Wi-Fi: %v", err)
		}

		linhas := strings.Split(string(out), "\n")

		var currentSSID string

		for _, linha := range linhas {
			linha = strings.TrimSpace(linha)

			if strings.HasPrefix(linha, "SSID") {
				partes := strings.SplitN(linha, ":", 2)
				if len(partes) > 1 {
					currentSSID = strings.TrimSpace(partes[1])
					if currentSSID == "" {
						currentSSID = "[Rede Oculta]"
					}
				}
			}

			if strings.HasPrefix(linha, "Sinal") || strings.HasPrefix(linha, "Signal") {
				partes := strings.SplitN(linha, ":", 2)
				if len(partes) > 1 {
					sinal := strings.TrimSpace(partes[1])
					fmt.Printf("%-25s | Sinal: %s\n", currentSSID, sinal)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(wifiCmd)
}