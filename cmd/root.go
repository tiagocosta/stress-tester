/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tiagocosta/stress-tester/internal/stress"
)

var stressTester stress.Tester

var rootCmd = &cobra.Command{
	Use:   "stress-tester",
	Short: "A tool for stress testing a service/URL",
	Run: func(cmd *cobra.Command, args []string) {
		stressTester.Stress()
		fmt.Printf("Quantidade total de requests realizadas: %d\n", stressTester.TotalRequests)
		for k, v := range stressTester.MapStatusCode {
			fmt.Printf("Quantidade de requests com status HTTP %d: %d\n", k, v)
		}
		fmt.Printf("Tempo total gasto na execução: %s\n", stressTester.TimeElapsed)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&stressTester.URL, "url", "u", "http://localhost:8080", "URL to be stressed")
	rootCmd.Flags().IntVarP(&stressTester.Requests, "requests", "r", 1000, "number of requsts")
	rootCmd.Flags().IntVarP(&stressTester.Concurrency, "concurrency", "c", 10, "number of concurrent calls to the URL")
}
