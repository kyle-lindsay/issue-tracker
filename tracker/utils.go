package tracker

import (
    "issue-tracker/model"
    
    "reflect"
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

func GetMatchingIssues(criteria model.IssueFilter) ([]model.Issue, error) {
    var issues []model.Issue
    var matchingIssues []model.Issue

    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    criteriaVal := reflect.ValueOf(criteria)
    match := false

    for i := range issues {
        match = true 
        issueVal := reflect.ValueOf(issues[i])

        for j:=0; j<criteriaVal.NumField(); j++ {
            fieldName := criteriaVal.Type().Field(j).Name

            if fieldName == "ActiveFilters" {
                continue
            }

            if !criteria.ActiveFilters[fieldName] {
                continue
            }

            criteriaField := criteriaVal.Field(j)
            cleanCriteria := reflect.Indirect(criteriaField)
            cleanIssue := reflect.Indirect(issueVal.Field(j))

            if criteriaField.Kind() == reflect.Ptr && criteriaField.IsNil() {
                continue
            }

            if !cleanCriteria.IsValid() {
                continue
            }

            criteriaFieldVal := cleanCriteria.Interface()
            issueFieldVal := cleanIssue.Interface()

            if fmt.Sprintf("%v", criteriaFieldVal) == "any" {
                continue
            }

            if !reflect.DeepEqual(criteriaFieldVal, issueFieldVal) {
                match = false
                break
            }
        }

        if !match {
            continue
        }
    
        matchingIssues = append(matchingIssues, issues[i])
    }
    return matchingIssues, nil
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
