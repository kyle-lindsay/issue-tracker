package model

import "time"

type Priority int
const (
    UnknownPriority Priority = iota
    Low
    Medium
    High
)
func (p Priority) String() string {
    strings := [...]string{"Unknown", "Low", "Medium", "High"}
    if p < Low || p > High {
        return "Unknown"
    }
    return strings[p]
}

type Status int
const (
    UnknownStatus Status = iota
    Open
    Closed
    Deleted
)
func (s Status) String() string {
    strings := [...]string{"Unknown", "Open", "Closed", "Deleted"}
    if s < Open || s > Deleted {
        return "Unknown"
    }
    return strings[s]
}

type Issue struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Priority Priority `json:"priority"`
    Status Status `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    LastUpdated time.Time `json:"last_updated"`
}
