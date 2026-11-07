package cmd

import (
    "fmt"
    "os"
    "text/tabwriter"
    "strings"

    "issue-tracker/model"
    "issue-tracker/tracker"
	"github.com/spf13/cobra"
)

var (
    listCount int
    listStatus string
    listPriority string
    oldestFirst bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists saved issues",
	Long: `List and view your project issues in a structured table.

By default this command displays a table of all saved issues with the [open] status, sorted by the creation date.
You can filter the results by specific priorities, statuses, or limit the total number of issues listed.`,
    Example:`    # List all open issues
        issues list
    # List all issues with the [high] priority
        issues list --status any --priority high
    # List the five oldest open issues
        issues list --limit 5 --oldest
    `,
	Run: func(cmd *cobra.Command, args []string) {
        var chosenStatus model.Status
        var chosenPriority model.Priority
        var printIssues []model.Issue
        
        listPriority = strings.ToLower(listPriority)
        if !IsValidQueryPriority(listPriority) {
            cmd.PrintErrf("Error: Invalid priority %q. Must be one of: %v\n", listPriority, GetValidQueryPriorities())
            return
        }
        
        listStatus = strings.ToLower(listStatus)
        if !IsValidQueryStatus(listStatus) {
            cmd.PrintErrf("Error: Invalid status %q. Must be one of: %v\n", listStatus, GetValidQueryStatuses())
            return
        }

        filter := model.IssueFilter{
            ActiveFilters: make(map[string]bool),
        }

        if listStatus != "" && listStatus != "any" {
            chosenStatus = ParseStatus(listStatus)
            filter.Status = &chosenStatus
            filter.ActiveFilters["Status"] = true
        }

        if listPriority != "" && listPriority != "any" {
            chosenPriority = ParsePriority(listPriority)
            filter.Priority = &chosenPriority
            filter.ActiveFilters["Priority"] = true
        }

        foundIssues, err := tracker.GetMatchingIssues(filter)
        if err != nil {
            cmd.PrintErrf("Error: Failed to list issues: %v\n", err)
            return
        }
        foundIssues = tracker.SortIssues(foundIssues, oldestFirst)
       
        foundIssuesLength := len(foundIssues)
        
        if foundIssuesLength == 0 {
            if chosenStatus == model.Open {
                fmt.Println("No issues to tackle, great work!")
                return
            } else {
                fmt.Println("No issues found")
                return
            }
        }
    
        if foundIssuesLength < listCount {
            listCount = foundIssuesLength
        }

        printIssues = foundIssues[:listCount]

        writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

        fmt.Fprintln(writer, "ID\tTITLE\tPRIORITY\tSTATUS\tCREATED AT\tLAST UPDATED")

        for _, issue := range printIssues {
            createdStr := issue.CreatedAt.Format("2006-01-02 15:04")
            updatedStr := issue.LastUpdated.Format("2006-01-02 15:04")


            fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\t%s\n",
                issue.Id,
                issue.Title,
                issue.Priority.String(),
                issue.Status.String(),
                createdStr,
                updatedStr,
            )
        }

        writer.Flush()
    },
}

func init() {
	rootCmd.AddCommand(listCmd)
    
    listCmd.Flags().IntVarP(&listCount, "limit", "l", 5, "The number of issues to list")
    listCmd.Flags().StringVarP(&listStatus, "status", "s", "open", "Limit to a certain status only")
    listCmd.Flags().StringVarP(&listPriority, "priority", "p", "any", "Limit to a certain priority only")
    listCmd.Flags().BoolVarP(&oldestFirst, "oldest", "o", false, "Sort issues by oldest first within each priority group")
}
