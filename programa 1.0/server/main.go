package main

import (
    "log"
    "net/http"
    "myapp/server/database"
    "myapp/server/handlers"
)

func main() {
    db, err := database.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/events", handlers.GetEvents(db))

    log.Println("Server starting on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}