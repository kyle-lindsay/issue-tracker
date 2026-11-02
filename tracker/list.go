package tracker

import (
    "issue-tracker/model"
    "sort"
)

func FindIssues(priority model.Priority, status model.Status) ([]model.Issue, error) {
    var issues []model.Issue
    var foundIssues []model.Issue
        
    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    for _, issue := range issues {
        if priority != model.UnknownPriority && issue.Priority != priority {
            continue
        }

        if status != model.UnknownStatus && issue.Status != status {
            continue
        }

        foundIssues = append(foundIssues, issue)
    }

    return foundIssues, nil
}

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
