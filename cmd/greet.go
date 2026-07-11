package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var upper bool

var greetCmd = &cobra.Command{
	Use:   "greet [nome]",
	Short: "Cumprimenta o usuário",
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		message := fmt.Sprintf("Fala, %s! Seja bem-vindo ao mundo Go!", name)

		if upper {
			message = strings.ToUpper(message)
		}

		fmt.Println(message)
	},
}

func init() {
	RootCmd.AddCommand(greetCmd)

	greetCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Exibe a mensagem em letras maiúsculas")
}