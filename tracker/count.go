package tracker

import(
    "issue-tracker/model"
)

func Count() (map[model.Status]map[model.Priority]int, error) {
    issues, err := GetIssues()
    if err != nil {
        return nil, err
    }

    var counts = make(map[model.Status]map[model.Priority]int)

    for _, issue := range issues {
        if issue.Status == model.UnknownStatus || issue.Priority == model.UnknownPriority {
            continue
        }

        if counts[issue.Status] == nil {
            counts[issue.Status] = make(map[model.Priority]int)
        }

        counts[issue.Status][issue.Priority] += 1
    }

    return counts, nil
}
