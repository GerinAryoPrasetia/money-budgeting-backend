package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "moneybudgeting-backend",
	Short: "Welcome to Money Budgeting Backend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Money Budgeting Backend")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
