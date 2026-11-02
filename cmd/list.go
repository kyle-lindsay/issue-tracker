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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists saved issues",
	Long: `Lists the first n issues in order of priority and time`,
	Run: func(cmd *cobra.Command, args []string) {
        var chosenStatus model.Status
        var chosenPriority model.Priority
        var printIssues []model.Issue
        
        listPriority = strings.ToLower(listPriority)
        if !IsValidQueryPriority(listPriority) {
            cmd.PrintErrf("Error: Invalid priority %q. Must be one of: %v\n", listPriority, GetValidPriorities())
            return
        }
        
        listStatus = strings.ToLower(listStatus)
        if !IsValidStatus(listStatus) {
            cmd.PrintErrf("Error: Invalid status %q. Must be one of: %v\n", listStatus, GetValidStatuses())
            return
        }

        chosenStatus = ParseStatus(listStatus)
        chosenPriority = ParsePriority(listPriority)

        foundIssues, err := tracker.FindIssues(chosenPriority, chosenStatus)
        if err != nil {
            cmd.PrintErrf("Error: Failed to list issues: %v\n", err)            
        }
        foundIssues = tracker.SortIssues(foundIssues, oldestFirst)
       
        foundIssuesLength := len(foundIssues)
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
    
    listCmd.Flags().IntVarP(&listCount, "limit", "n", 5, "The number of issues to list")
    listCmd.Flags().StringVarP(&listStatus, "status", "s", "open", "Limit to a certain status only")
    listCmd.Flags().StringVarP(&listPriority, "priority", "p", "any", "Limit to a certain priority only")
    listCmd.Flags().BoolVarP(&oldestFirst, "oldest", "o", false, "Sort issues by oldest first within each priority group")
}
