package cmd

import (
	"cep_desafio/app/services"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BuscarCEP",
	Short: "Busca informações de um CEP em dois serviços diferentes",
	Long:  `Este comando busca informações de um CEP utilizando os serviços ViaCEP e BrasilAPI.`,
	Run: func(cmd *cobra.Command, args []string) {
		waitgroup := sync.WaitGroup{}
		waitgroup.Add(3)
		cep, _ := strconv.ParseInt(args[1], 10, 64)

		ch := make(chan interface{})
		done := make(chan bool)
		go services.GetCEPInfo(cep, "viacep", ch, &waitgroup)
		go services.GetCEPInfo(cep, "brasilapi", ch, &waitgroup)
		go Spinner(100*time.Millisecond, done)

		select {
		case result := <-ch:
			done <- true
			resultJsonFormatted, err := json.MarshalIndent(result, "", "  ")
			if err == nil {
				fmt.Printf("\rO resultado da busca é:\n%s\n", resultJsonFormatted)
			}
			return
		case <-time.After(1 * time.Second):
			fmt.Println("\rA busca demorou mais de 1 segundo e foi cancelada.")
		}

		waitgroup.Wait()
	}}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar o comando: %v\n", err)
		os.Exit(1)
	}
}

func Spinner(delay time.Duration, stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Print("\r")
			return
		default:
			for _, r := range `-\|/` {
				select {
				case <-stopChan:
					fmt.Print("\r")
					return
				default:
					fmt.Printf(" Buscando o CEP  \r%c", r)
					time.Sleep(delay)
				}
			}
		}
	}
}
