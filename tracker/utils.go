package tracker

import (
    "issue-tracker/model"
    "os"
    "strings"
    "encoding/json"
    "fmt"
)

const jsonFile = ".issues.json"

func GetIssues() ([]model.Issue, error) {
    var issues []model.Issue

    if _, err := os.Stat(jsonFile); err == nil {
        data, err := os.ReadFile(jsonFile)
        if err != nil {
            return nil, err
        }

        _ = json.Unmarshal(data, &issues)
    }

    
    return issues, nil
}

func GetIssue(id int) (*model.Issue, error) {
    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    for i := range issues {
        if issues[i].Id == id {
            return &issues[i], nil
        }
    }

    return nil, nil
}

func GetIssuesWithStatus(status model.Status) ([]model.Issue, error) {
    var issues []model.Issue
    var issuesWithStatus []model.Issue

    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    for _, issue := range issues {
        if issue.Status == status {
            issuesWithStatus = append(issuesWithStatus, issue)
        }
    }
    return issuesWithStatus, nil
}

func UpdateGitIgnore() error {
    if _, err := os.Stat(".gitignore"); !os.IsNotExist(err) {
        data, err := os.ReadFile(".gitignore")
        if err != nil {
            return err
        }

        if strings.Contains(string(data), ".issues.json") {
            return nil
        }

        newData := fmt.Sprintf("%s\n.issues.json\n", string(data))
        return os.WriteFile(".gitignore", []byte(newData), 0644)
    }

    return nil
}
