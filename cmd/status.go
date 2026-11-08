package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays information abouut your tracked issues",
	Long: `Counts the number of issues, seperated by priority within each status`,
    Example: `    # Display tracked issue status
        issues status`,
	Run: func(cmd *cobra.Command, args []string) {
        counts, err := tracker.Count()
        if err != nil {
            cmd.PrintErrf("Error: could not display status: ", err)
        }

        statuses := []model.Status{model.Open, model.Closed, model.Deleted}
        priorities := []model.Priority{model.High, model.Medium, model.Low}

        totalCount := 0
        for _, status := range statuses {
            fmt.Printf("%s:\n", status.String())
            issueCount := 0
            for _, priority := range priorities {
                fmt.Printf("    %s: %d\n", priority.String(), counts[status][priority])
                issueCount += counts[status][priority]
            }
            fmt.Printf("    Total: %d\n", issueCount)
            fmt.Printf("\n")
            totalCount += issueCount
        }
        fmt.Printf("Total number of issues: %d\n", totalCount)
    },
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
