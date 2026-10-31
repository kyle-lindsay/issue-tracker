package cmd

import (
    "issue-tracker/model"  
)

func ParsePriority(priority string) model.Priority {
    var chosenPriority model.Priority
        switch priority {
            case "low":
                chosenPriority = model.Low
            case "high":
                chosenPriority = model.High
            default:
                chosenPriority = model.Medium
        }
    return chosenPriority
}
