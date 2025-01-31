package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Отметить задачу как невыполненная: todo [номер]",
	Run: func(cmd *cobra.Command, args []string) {
		index := parseIndex(args[0])

		if index == -1 || index > (len(TaskList)+1) {
			fmt.Println("Error, index out of range")
			return
		}

		TaskList[index].Status = "✗"

		err := saveTasks()
		if err != nil {
			fmt.Println("Error, saving tasks ", err)
			return
		}
		fmt.Printf("Task %d marked as todo!\n", index+1)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

}
