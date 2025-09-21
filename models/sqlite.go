package models

import {
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
}

db, err := sql.Open("sqlite3", "./users.db")
if err != nil {
	panic(err)
}
defer db.Close()

_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
                  id TEST PRIMARY KEY,
                  password TEXT
)`)

_, err = db.Exec("INSERT INTO users(id, password)
                  VALUES (?, ?)", id, password)

row := db.QueryRow("SELECT password FROM users WHERE id = ?", id)
var dbPassword string
err = row.Scan(&dbPassword)
if err == nil && dbPassword == password {
	
}
