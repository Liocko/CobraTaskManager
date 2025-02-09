package cmd

import (
	"database/sql"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "cobraTaskManager",
	Short: "Simple cobra task manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal("Error opening database: ", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal("Error pinging database: ", err)
		}
		if err := createTables(db); err != nil {
			log.Fatal("Error creating tables: ", err)
		}

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			db.Close()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
