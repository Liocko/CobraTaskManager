/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить задачу в формате: Название | Описание",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		parts := strings.SplitN(input, " | ", 2)
		if len(parts) != 2 {
			fmt.Println("usage: add [name] [description]")
		}
		title := strings.TrimSpace(parts[0])
		description := strings.TrimSpace(parts[1])
		newTask := Task{
			Title:       title,
			Description: description,
			Status:      "✗",
		}
		TaskList = append(TaskList, newTask)

		if err := saveTasks(); err != nil {
			fmt.Println("Error saving tasks: " + err.Error())
			return
		}
		fmt.Printf("Task added: %s\n", newTask.Title)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
