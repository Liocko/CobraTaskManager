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

		if index == -1 || index > (len(TaskList)+1) {
			fmt.Println("Error, index out of range")
			return
		}

		TaskList[index].Status = "✓"

		err := saveTasks()
		if err != nil {
			fmt.Println("Error, saving tasks ", err)
			return
		}
		fmt.Printf("Task %d marked as done!\n", index+1)

	},
}

func parseIndex(input string) int {
	var index int
	n, err := fmt.Sscanf(input, "%d", &index)
	if n != 1 || err != nil {
		fmt.Println("Error, type a digit ", err)
		return -1
	}
	return index - 1
}
func init() {
	rootCmd.AddCommand(doneCmd)

}
