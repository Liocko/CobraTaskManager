package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Отметить задачу как выполненная: done [номер]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := parseIndex(args[0])
		result, err := db.Exec(`UPDATE tasks SET status = "✓" WHERE id = ?`, index+1)
		if err != nil {
			fmt.Println("Error marking task as done: ", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting rows affected: ", err)
		}
		if rowsAffected == 0 {
			fmt.Printf("Task %d not found.\n", index+1)
		} else {
			fmt.Printf("Task %d marked as done!\n", index+1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
