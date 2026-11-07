package cmd

import (
    "strconv"
	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

var reopenCmd = &cobra.Command{
	Use:   "reopen [id]",
	Short: "Set this closed issue's status to [Open]",
	Long: `Updates the status of the issue to [Open]. Note, cannot repoen deleted issues, you must use 'issues recover' first.`,
    Example: `    # Reopen a specific closed issue by id
        issues open 42`,
    DisableFlagsInUseLine: true,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        var issueId, err = strconv.Atoi(args[0])
        if err != nil {
            cmd.PrintErrf("Error: invalid issue id\n")
            return
        }
        issue, err := tracker.SetOpenClosed(issueId, model.Open)
        if err != nil {
            cmd.PrintErrf("Error: failed to reopen issue: %v\n", err)
            return
        }

        cmd.Printf("Reopened issue #%d (%s)\n", issue.Id, issue.Title)
    },
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
