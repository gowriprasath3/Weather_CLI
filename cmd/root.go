package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weathercli",
	Short: "app to get the current weather",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root cmd...")
	},
}

func Execute() {
	log.Println("execute root cmd...")
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	log.Println("init root cmd...")
	rootCmd.AddCommand(currentCmd)
}
