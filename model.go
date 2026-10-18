package main

import "time"

type Priority int
const (
    Low Priority = iota
    Medium
    High
)
func (p Priority) String() string {
    strings := [...]string{"Low", "Medium", "High"}
    if p < Low || p > High {
        return "Unknown"
    }
    return strings[p]
}

type Status int
const (
    Open Status = iota
    Closed
    Deleted
)
func (s Status) String() string {
    strings := [...]string{"Open", "Closed", "Deleted"}
    if s < Open || s > Deleted {
        return "Unknown"
    }
    return strings[s]
}

type Issue struct {
    Id int `json:id`
    Title string `json:title`
    Description string `json:description`
    Priority Priority `json:priority`
    Status Status `json:status`
    CreatedAt time.Time `json:created_at`
}
