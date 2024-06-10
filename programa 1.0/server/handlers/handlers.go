package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
)

type Event struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Location    string `json:"location"`
    Datetime    string `json:"datetime"`
}

func GetEvents(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, title, description, location, datetime FROM events")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var events []Event
        for rows.Next() {
            var event Event
            if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Datetime); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            events = append(events, event)
        }
        if err := rows.Err(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(events); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}