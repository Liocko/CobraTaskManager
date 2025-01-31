/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

// tasklistCmd represents the TaskList command
var tasklistCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывести все задачи в виде таблицы: list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(TaskList) == 0 {
			fmt.Println("Task list is empty.")
			return
		}
		//	fmt.Println("List of tasks:")
		//	for i, task := range TaskList {
		//		fmt.Printf("%d. %s - %s [%s]\n", i+1, task.Title, task.Description, task.Status)
		//	}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Task name", "Task description", "Done"})
		for i, task := range TaskList {
			t.AppendRow(table.Row{i + 1, task.Title, task.Description, task.Status})
			t.AppendSeparator()
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(tasklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasklistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
