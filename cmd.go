package main

import (
	_ "aisecurity/ent/dao/runtime"
	"aisecurity/utils"
	"aisecurity/utils/db"
	zap2 "aisecurity/utils/log/zap"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aisecurity",
	Short: "A CLI tool for ent migrations and generation",
	Run: func(cmd *cobra.Command, args []string) {
		// Display help if no command is provided
		fmt.Println("Please provide a command. Use 'ent --help' for usage.")
	},
}

// Execute is the entry point for the CLI
func execute() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "cmd1",
		Short: "short cmd1",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "cmd2",
		Short: "short cmd2",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// if rebuild cmd bin, rename the function main2 to main
func main2() {
	InitMain()
	defer func(Logger *zap2.Logger) {
		err := Logger.Sync()
		if err != nil {
			log.Fatalf("failed syncing logger: %v", err)
		}
	}(utils.Logger)

	defer func() {
		err := db.EntClient.Close()
		if err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()

	db.Generate()
	db.Migrate()

	execute()
}
