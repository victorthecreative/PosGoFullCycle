package cmd

import (
	"cep_desafio/app/services"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BuscarCEP",
	Short: "Busca informações de um CEP em dois serviços diferentes",
	Long:  `Este comando busca informações de um CEP utilizando os serviços ViaCEP e BrasilAPI.`,
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cep, _ := strconv.ParseInt(args[1], 10, 64)

		ch := make(chan interface{})
		go services.GetCEPInfo(cep, "viacep", ch)
		go services.GetCEPInfo(cep, "brasilapi", ch)

		select {
		case result := <-ch:
			resultJsonFormatted, err := json.MarshalIndent(result, "", "  ")
			if err == nil {
				fmt.Printf("\rO resultado da busca é:\n%s\n", resultJsonFormatted)
			}
		case <-time.After(1 * time.Second):
			fmt.Println("\rA busca demorou mais de 1 segundo e foi cancelada.")
		}
	}}

func Execute() {
	time.Sleep(5 * time.Second)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar o comando: %v\n", err)
		os.Exit(1)
	}
}
