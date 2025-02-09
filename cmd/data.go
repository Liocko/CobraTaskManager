package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func parseIndex(input string) int {
	var index int
	n, err := fmt.Sscanf(input, "%d", &index)
	if n != 1 || err != nil {
		fmt.Println("Error, type a digit ", err)
		return -1
	}
	return index - 1
}

func createTables(db *sql.DB) error {
	createTable := `CREATE TABLE IF NOT EXISTS tasks (
  		id INTEGER PRIMARY KEY AUTOINCREMENT,
  		title VARCHAR(255) NOT NULL,
    	description TEXT NOT NULL,
    	status INTEGER NOT NULL
);`
	_, err := db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}

func addListToDb(task Task) {
	for _, task := range TaskList {
		_, err := db.Exec("INSERT INTO tasks(title, description, status) VALUES(?, ?, ?)", task.Title, task.Description, task.Status)
		if err != nil {
			log.Fatal(err)
		}
	}
}
