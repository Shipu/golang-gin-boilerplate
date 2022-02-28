package main

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "command",
	Short: "A simple generate list",
	Long:  `A simple generate list`,
	Run:   hello,
}

func hello(cmd *cobra.Command, args []string) {
	log.Println("HELLO")
}

func init() {
	rootCmd.AddCommand(cliCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
