package cmd

import (
    "strconv"
	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

var closeCmd = &cobra.Command{
	Use:   "close [id]",
	Short: "Set this issue's status to [Closed]",
	Long: `Updates the status of the issue to [Closed]`,
    Example: `    # Close a specific issue by id
        issues close 42`,
    DisableFlagsInUseLine: true,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        var issueId, err = strconv.Atoi(args[0])
        if err != nil {
            cmd.PrintErrf("Error: invalid issue id\n")
            return
        }
        err = tracker.SetOpenClosed(issueId, model.Closed)
        if err != nil {
            cmd.PrintErrf("Error: failed to close issue: %v\n", err)
            return
        }

        cmd.Printf("Closed issue #%d\n", issueId)
    },
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
