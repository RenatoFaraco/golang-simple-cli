package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "golang-simple-cli",
	Short: "Uma CLI interativa em Go",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		
		fmt.Println("==================================================")
		fmt.Println("       BEM-VINDO À SUA CLI INTERATIVA EM GO       ")
		fmt.Println(" Digite 'help' para comandos ou 'exit' para sair. ")
		fmt.Println("==================================================")

		for {
			fmt.Print("golang-cli> ")
			
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Erro ao ler comando:", err)
				continue
			}

			input = strings.TrimSpace(input)

			if input == "" {
				continue
			}

			if input == "exit" || input == "quit" {
				fmt.Println("Até logo!")
				break
			}

			cmdArgs := strings.Fields(input)
			
			cmd.SetArgs(cmdArgs)
			
			if err := cmd.Execute(); err != nil {
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