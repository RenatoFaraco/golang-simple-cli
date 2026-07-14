package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "Lista as redes Wi-Fi disponíveis e a intensidade do sinal",
	Long:  `Executa o utilitário do Windows para escanear e listar os SSIDs e sinais das redes próximas.`,
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan).SprintFunc()
		whiteBold := color.New(color.FgWhite, color.Bold).SprintFunc()

		fmt.Println(cyan("🔎 Escaneando redes Wi-Fi próximas..."))
		fmt.Println(cyan("--------------------------------------------------"))

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
					sinalRaw := strings.TrimSpace(partes[1]) // Ex: "92%"
					
					// Remove o símbolo de % e espaços para converter para número
					sinalLimpo := strings.ReplaceAll(sinalRaw, "%", "")
					sinalLimpo = strings.TrimSpace(sinalLimpo)
					
					// Converte o sinal em inteiro para aplicar a cor correspondente
					sinalNum, err := strconv.Atoi(sinalLimpo)
					
					var sinalColorido string
					if err != nil {
						// Fallback caso falhe a conversão
						sinalColorido = color.YellowString(sinalRaw)
					} else {
						// Lógica de cores do sinal
						if sinalNum >= 75 {
							sinalColorido = color.GreenString(sinalRaw)
						} else if sinalNum >= 40 {
							sinalColorido = color.YellowString(sinalRaw)
						} else {
							sinalColorido = color.RedString(sinalRaw)
						}
					}

					// Printa o SSID com destaque e o sinal colorido correspondente
					fmt.Printf("📶 %-25s | Sinal: %s\n", whiteBold(currentSSID), sinalColorido)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(wifiCmd)
}