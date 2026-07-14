package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "golang-simple-cli",
	Short: "Uma CLI interativa em Go",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		
		// Criando estilos reutilizáveis
		cyan := color.New(color.FgCyan).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		blueBold := color.New(color.FgBlue, color.Bold).SprintFunc()

		fmt.Println(cyan("=================================================="))
		fmt.Printf("       %s       \n", blueBold("BEM-VINDO À SUA CLI INTERATIVA EM GO"))
		fmt.Printf(" Digite '%s' para comandos ou '%s' para sair. \n", yellow("help"), yellow("exit"))
		fmt.Println(cyan("=================================================="))

		for {
			color.New(color.FgGreen, color.Bold).Print("golang-cli> ")
			
			input, err := reader.ReadString('\n')
			if err != nil {
				color.Red("Erro ao ler comando: %v", err)
				continue
			}

			input = strings.TrimSpace(input)

			if input == "" {
				continue
			}

			if input == "exit" || input == "quit" {
				fmt.Println(green("Até logo! 👋"))
				break
			}

			cmdArgs := strings.Fields(input)
			
			cmd.SetArgs(cmdArgs)
			
			if err := cmd.Execute(); err != nil {
				// Erro tratado pelo Cobra
			}
			fmt.Println() 
		}
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}