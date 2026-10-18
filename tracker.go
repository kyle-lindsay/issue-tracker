package main

import (
    "errors"
    "os"
    "encoding/json"
)

const jsonFile = ".issues.json"

func createIssue(newIssue Issue) error {
    if newIssue.Title == "" {
        return errors.New("issue title cannot be empty")
    }

    var issues []Issue

    if _, err := os.Stat(jsonFile); err == nil {
        data, err := os.ReadFile(jsonFile)
        if err != nil {
                return err
        }

        _ = json.Unmarshal(data, &issues)
    }

    maxId := 0
    for _, existingIssue := range issues {
        if existingIssue.Title == newIssue.Title {
            return errors.New("issue title must be unique")
        }
        if existingIssue.Id > maxId {
            maxId = existingIssue.Id
        }
    }

    newIssue.Id = maxId + 1

    newIssue.Status = Open

    issues = append(issues, newIssue)
    jsonData, err := json.MarshalIndent(issues, "", " ")
    if err != nil {
        return err
    }

    return os.WriteFile(jsonFile, jsonData, 0644)
}
