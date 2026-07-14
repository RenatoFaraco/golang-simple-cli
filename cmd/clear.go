package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Limpa a tela do terminal",
	Long:  `Executa o comando de limpeza de tela compatível com o sistema operacional atual.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Tenta limpar a tela rodando o comando nativo do sistema operacional
		var clearCmd *exec.Cmd

		if runtime.GOOS == "windows" {
			clearCmd = exec.Command("cmd", "/c", "cls")
		} else {
			// Como estamos rodando no WSL (Linux), ele vai cair aqui
			clearCmd = exec.Command("clear")
		}

		// Vincula a saída do comando de limpar diretamente ao terminal atual
		clearCmd.Stdout = os.Stdout
		err := clearCmd.Run()

		if err != nil {
			// Se o comando do sistema falhar por algum motivo, usamos o plano B:
			// Código ANSI de escape para limpar a tela e mover o cursor para o topo
			fmt.Print("\033[H\033[2J")
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}