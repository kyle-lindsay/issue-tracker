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
