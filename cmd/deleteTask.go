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
		if index == -1 || index > (len(TaskList)+1) {
			fmt.Println("Error, index out of range")
			return
		}
		TaskList = append(TaskList[:index], TaskList[index+1:]...)
		err := saveTasks()
		if err != nil {
			fmt.Println("Error saving tasks", err)
			return
		}
		fmt.Printf("Task %d has been deleted succesfully", index+1)
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
