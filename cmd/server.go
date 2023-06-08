package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/GerinAryoPrasetia/money-budgeting-backend/internal/pkg/env"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "runs server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := env.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// _, err = sql.Open(config.PostgresDriver, config.PostgresSource)
	// if err != nil {
	// 	log.Fatalf("could not connect to postgres database: %v", err)
	// }
	fmt.Println("server running...")
}
