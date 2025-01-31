package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Вывод текущего времени: time",
	Run: func(cmd *cobra.Command, args []string) {
		location, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			fmt.Println("Error loading location", err)
			return
		}
		currentTime := time.Now().In(location)
		fmt.Printf("Current time is %s", currentTime.Format("2006-01-02 15:04:05 MST"))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("location", "l", "Europe/Moscow", "Location")

}
