package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobraTaskManager",
	Short: "Simple cobra task manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	loadTasks()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
