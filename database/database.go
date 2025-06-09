package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id    int
	Title string
}

var DB *sql.DB
var init_db_statement string = `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`
var read_db_statement string = "SELECT * from todos"
var insert_into_db_statement string = `INSERT INTO todos(title) VALUES(?)`
var delete_from_db_statement string = `DELETE FROM todos WHERE id=?`

func InitDB() {
	var err error
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to find home directory:", err)
	}
	dbPath := filepath.Join(homeDir, ".todo", "data.db")
	err = os.MkdirAll(filepath.Dir(dbPath), 0755)
	if err != nil {
		log.Fatalf("Failed to create config directory: %v", err)
	}
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(init_db_statement)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, init_db_statement)
	}
}

func ReadDB() []Todo {
	rows, err := DB.Query(read_db_statement)
	if err != nil {
		log.Fatal("error getting data from database")
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Title); err != nil {
			log.Fatal("Error reading next row from db")
		}
		todos = append(todos, todo)
	}
	return todos
}

func InsertIntoDB(data string) {
	if _, err := DB.Exec(insert_into_db_statement, data); err != nil {
		log.Fatalf("Error inserting data into db %v", err)
	}
}

func DeleteFromDB(id int) {
	if _, err := DB.Exec(delete_from_db_statement, id); err != nil {
		log.Fatalf("Error deleting data from db: %v", err)
	}
}
