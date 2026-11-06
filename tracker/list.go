package tracker

import (
    "issue-tracker/model"
    "sort"
)

func SortIssues(issues []model.Issue, oldestFirst bool) []model.Issue {
    sort.Slice(issues, func(i, j int) bool {
        if issues[i].Priority != issues[j].Priority {
            return issues[i].Priority > issues[j].Priority
        }

        if oldestFirst {
            return issues[i].CreatedAt.Before(issues[j].CreatedAt)
        }

        return issues[i].CreatedAt.After(issues[j].CreatedAt)
    })
    return issues
}
