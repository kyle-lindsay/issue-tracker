package tracker

import (
    "os"
    "reflect"
    "fmt"
    "time"

    "encoding/json"
    "issue-tracker/model"
)

func UpdateIssue(id int, updates model.Issue) (*model.Issue, error) {
    var issues []model.Issue
    var updatedIssue model.Issue

    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    for i := range issues {
        if issues[i].Id == id {
            srcVal := reflect.ValueOf(updates)
            destVal := reflect.ValueOf(&issues[i]).Elem()

            for j:=0; j<srcVal.NumField(); j++ {
                srcFieldVal := srcVal.Field(j)
                destFieldVal := destVal.Field(j)

                if !destFieldVal.CanSet() {
                    continue
                }

                if !srcFieldVal.IsZero() {
                    destFieldVal.Set(srcFieldVal)
                }
            }
            
            issues[i].LastUpdated = time.Now()
            updatedIssue = issues[i]
            break
        }
    }


    if updatedIssue == (model.Issue{}) {
        return nil, fmt.Errorf("issue with id %d not found", id)
    }

    jsonData, err := json.MarshalIndent(issues, "", " ")
    if err != nil {
        return nil, err
    }

    return &updatedIssue, os.WriteFile(jsonFile, jsonData, 0644)
}

func SetStatus(id int, status model.Status) (*model.Issue, error) {
    issue, err := GetIssue(id)
    if err != nil {
        return nil, err
    }

    if issue.Status == status {
        return nil, fmt.Errorf("issue %d already has this status", id)
    }

    return UpdateIssue(id, model.Issue{Status: status})
}
