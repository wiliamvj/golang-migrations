package main

import (
  "database/sql"
  "fmt"
  "log"
  "os"

  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
)

func main() {
  // load .env file
  godotenv.Load()

  postgresURI := os.Getenv("DATABASE_URL")
  db, err := sql.Open("postgres", postgresURI)
  if err != nil {
    log.Panic(err)
  }
  err = db.Ping()
  if err != nil {
    db.Close()
    log.Panic(err)
  }

  fmt.Println("Connected to database")

  // keep the program running
  select {}
}
