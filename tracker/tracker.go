package tracker

import (
    "errors"
    "os"
    "fmt"
    "strings"
    "encoding/json"
    "issue-tracker/model"
)

func CreateIssue(newIssue model.Issue) error {
    if newIssue.Title == "" {
        return errors.New("issue title cannot be empty")
    }

    var issues []model.Issue

    if _, err := os.Stat(jsonFile); err == nil {
        data, err := os.ReadFile(jsonFile)
        if err != nil {
                return err
        }

        _ = json.Unmarshal(data, &issues)
    }

    maxId := 0
    for _, existingIssue := range issues {
        if existingIssue.Id > maxId {
            maxId = existingIssue.Id
        }
    }

    newIssue.Id = maxId + 1

    newIssue.Status = model.Open

    issues = append(issues, newIssue)
    jsonData, err := json.MarshalIndent(issues, "", " ")
    if err != nil {
        return err
    }

    return os.WriteFile(jsonFile, jsonData, 0644)
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
