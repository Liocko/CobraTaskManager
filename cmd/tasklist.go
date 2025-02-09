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
		rows, err := db.Query("SELECT id, title, description, status FROM tasks")
		if err != nil {
			fmt.Println("Error database query: ", err)
		}
		defer rows.Close()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Task name", "Task description", "Done"})

		var found bool
		for rows.Next() {
			var id int
			var title string
			var description string
			var status string
			if err := rows.Scan(&id, &title, &description, &status); err != nil {
				fmt.Println("Error reading database: ", err)
				return
			}
			t.AppendRow(table.Row{id, title, description, status})
			t.AppendSeparator()
		}
		if !found {
			fmt.Println("No tasks found")
		}
		t.SetStyle(table.StyleLight)
		t.Render()
	},
}

//		if err != nil {
//			fmt.Println(err)
//			log.Fatal(err)
//		}
//		fmt.Println(rows)
//
//		t.AppendRow(table.Row{i + 1, task.Title, task.Description, task.Status})
//		t.AppendSeparator()
//	}
//	t.SetStyle(table.StyleLight)
//	t.Render()
//},

func init() {
	rootCmd.AddCommand(tasklistCmd)

}
