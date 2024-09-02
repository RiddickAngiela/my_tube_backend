package config

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dsn string) (*sql.DB, error) {
    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    // Ensure the database connection is available
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error connecting to the database: %v", err)
    }

    return db, nil
}