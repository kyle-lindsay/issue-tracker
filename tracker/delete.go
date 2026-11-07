package tracker

import (
    "fmt"
    "encoding/json"
    "os"

    "issue-tracker/model"
)

func HardDelete(issueId int) error {
    issues, err := GetIssues()
    if err != nil {
        return err
    }

    var remainingIssues []model.Issue
    found := false

    for _, issue := range issues {
        if issue.Id == issueId {
            found = true
            continue
        }
        remainingIssues = append(remainingIssues, issue)
    }

    if !found {
        return fmt.Errorf("issue with id %d does not exist", issueId)
    }

    jsonData, err := json.MarshalIndent(remainingIssues, "", " ")
    if err != nil {
        return err
    }

    return os.WriteFile(jsonFile, jsonData, 0644)
}

func SoftDelete(issueId int) error {
    issue, err := GetIssue(issueId)
    if err != nil {
        return err
    }

    if issue == nil {
        return fmt.Errorf("issue with id %d does not exist", issueId)
    }

    if issue.Status == model.Deleted {
       return fmt.Errorf("issue with id %d has already been deleted", issueId)
    }

    _, err = UpdateIssue(issueId, model.Issue{Status: model.Deleted})
    return err
}

func RecoverSoftDelete(issueId int, status model.Status) (*model.Issue, error) {
    issue, err := GetIssue(issueId)
    if err != nil {
        return nil, err
    }

    if issue == nil {
        return nil, fmt.Errorf("issue with id %d does not exist", issueId)
    }

    if issue.Status != model.Deleted {
       return nil, fmt.Errorf("issue with id %d is not deleted", issueId)
    }

    return UpdateIssue(issueId, model.Issue{Status: status})
    
}
