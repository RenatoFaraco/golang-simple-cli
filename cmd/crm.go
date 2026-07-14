package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Comando pai 'crm'
var crmCmd = &cobra.Command{
	Use:   "crm",
	Short: "Painel administrativo para o GopherCRM",
	Long:  `Permite gerenciar leads, clientes e visualizar o status do servidor do GopherCRM direto pelo terminal.`,
}

// Subcomando 'crm login'
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Autentica o administrador no GopherCRM",
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan).SprintFunc()
		whiteBold := color.New(color.FgWhite, color.Bold).SprintFunc()

		fmt.Println(cyan("🔐 Iniciando autenticação no GopherCRM..."))
		fmt.Println(cyan("--------------------------------------------------"))

		// Simulação de requisição para a API do CRM
		fmt.Print("Enviando credenciais para o servidor... ")
		time.Sleep(1 * time.Second) // Dá aquele efeito de carregamento real

		fmt.Println(color.GreenString("Sucesso! ✅"))
		fmt.Printf("👤 %-12s %s\n", whiteBold("Usuário:"), color.YellowString("exemplo"))
		fmt.Printf("🔑 %-12s %s\n", whiteBold("Token JWT:"), color.BlueString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."))
		fmt.Println(cyan("--------------------------------------------------"))
		fmt.Println(color.GreenString("Sessão iniciada com sucesso. Você já pode gerenciar o CRM!"))
	},
}

func init() {
	// Adiciona o subcomando 'login' ao comando pai 'crm'
	crmCmd.AddCommand(loginCmd)

	// Adiciona o comando pai 'crm' à nossa CLI principal
	RootCmd.AddCommand(crmCmd)
}