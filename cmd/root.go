package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "simple-cli",
	Short: "Simple-CLI é uma ferramenta de exemplo construída em Go",
	Long:  "Uma CLI super veloz e leve criada para aprender os conceitos básicos de Golang.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Olá! Use '--help' ou '-h' para ver os comandos disponíveis.")
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