package cmd

import (
    "slices"

    "issue-tracker/model"  
)

func ParsePriority(priority string) model.Priority {
    var chosenPriority model.Priority
        switch priority {
            case "low":
                chosenPriority = model.Low
            case "medium":
                chosenPriority = model.Medium
            case "high":
                chosenPriority = model.High
            default:
                chosenPriority = model.UnknownPriority
        }
    return chosenPriority
}

func ParseStatus(status string) model.Status {
    var chosenStatus model.Status
    switch status {
        case "open":
            chosenStatus = model.Open
        case "closed":
            chosenStatus = model.Closed
        case "deleted":
            chosenStatus = model.Deleted
        default:
            chosenStatus = model.UnknownStatus
    }
    return chosenStatus
}

func IsValidPriority(priority string) bool {
    return slices.Contains(GetValidPriorities(), priority)
}

func IsValidQueryPriority(priority string) bool {
    return slices.Contains(getQueryPriorities(), priority)
}

func GetValidPriorities() []string {
    validPriorities := []string{"low", "medium", "high"}
    return validPriorities
}

func getQueryPriorities() []string {
    return append(GetValidPriorities(), "any")
}

func IsValidStatus(status string) bool {
    return slices.Contains(GetValidStatuses(), status)
}

func GetValidStatuses() []string {
    validStatuses := []string{"open", "closed", "deleted"}
    return validStatuses
}
