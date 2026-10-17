package main

import 'time'

type Priority int
const (
    Low Priority = iota
    Medium
    High
)
func (p Priority) String() string {
    strings := [...]string{'Low', 'Medium', 'High'}
    if p < Low || p > High {
        return 'Unknown'
    }
    return strings[p]
}

type Issue struct {
    Id int `json:id`
    Title string `json:title`
    Description string `json:description`
    Priority Priority `json:priority`
    CreatedAt time.Time `json:created_at`
}
