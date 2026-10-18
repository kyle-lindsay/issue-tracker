package cmd

import (
//	"fmt"
    "strconv"

	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

const jsonFile = ".issues.json"

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close [title]",
	Short: "Set this issue to [Closed]",
	Long: `Closes the issue`,
    DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
        var issueId, err = strconv.Atoi(args[0])
        if err != nil {
            cmd.PrintErrf("Error: Invalid Issue ID")
            return
        }
        issue, err := tracker.SetStatus(issueId, model.Closed)
        if err != nil {
            cmd.PrintErrf("Error: Failed to close issue: %v\n", err)
            return
        }

        cmd.Printf("Closed issue #%d (%s)\n", issue.Id, issue.Title)
    },
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
