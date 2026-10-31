package tracker

import (
    "os"
    "encoding/json"
    "issue-tracker/model"
)

func SetStatus(id int, status model.Status) (*model.Issue, error) {
    var issues []model.Issue
    var updatedIssue model.Issue

    if _, err := os.Stat(jsonFile); err == nil {
        data, err := os.ReadFile(jsonFile)
        if err != nil {
            return nil, err
        }

        _ = json.Unmarshal(data, &issues)
    }

    for i, issue := range issues {
        if issue.Id == id {
            issues[i].Status = status
            updatedIssue = issue
            break
        }
    }

    jsonData, err := json.MarshalIndent(issues, "", " ")
    if err != nil {
        return nil, err
    }

    return &updatedIssue, os.WriteFile(jsonFile, jsonData, 0644)
}
