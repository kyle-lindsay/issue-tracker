package tracker

import (
    "issue-tracker/model"
    "os"
    "encoding/json"
)

const jsonFile = ".issues.json"

func GetIssuesWithStatus(status model.Status) ([]model.Issue, error) {
    var issues []model.Issue

    if _, err := os.Stat(jsonFile); err == nil {
        data, err := os.ReadFile(jsonFile)
        if err != nil {
            return nil, err
        }

        _ = json.Unmarshal(data, &issues)
    }

    var openIssues []model.Issue
    for _, currentIssue := range issues {
        if currentIssue.Status == status {
            openIssues = append(openIssues, currentIssue)
        }
    }
    
    return openIssues, nil
}

