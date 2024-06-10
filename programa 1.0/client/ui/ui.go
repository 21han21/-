package ui

import (
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

func FetchEvents() ([]Event, error) {
    resp, err := http.Get("http://localhost:8081/events")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var events []Event
    if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
        return nil, err
    }
    return events, nil
}