/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// deleteTaskCmd represents the deleteTask command
var deleteTaskCmd = &cobra.Command{
	Use:   "remove",
	Short: "Удалить задачу по номеру: remove [номер]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := parseIndex(args[0])
		result, err := db.Exec("DELETE FROM tasks WHERE id = ?", index+1)
		if err != nil {
			fmt.Println("Error deleting task: ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting rows affected: ", err)
		}
		if rowsAffected == 0 {
			fmt.Printf("Task %d not found.\n", index+1)
		} else {
			fmt.Printf("Task %d deleted successfully.\n", index+1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)

}
