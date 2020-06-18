package main

import (
	"database/sql"
	"fmt"
	projectDB "github.com/grandima/go-task-management-app/db/storage"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

func main() {

	connStr := "user=grandima dbname=task_management_app sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var storage = projectDB.NewProjectService(db)

	fmt.Println(storage.Index())
}
