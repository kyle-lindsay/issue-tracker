package cmd

import (
	"fmt"
    "strconv"
    "issue-tracker/tracker"

	"github.com/spf13/cobra"
)

var hardDelete bool

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete an issue entry",
	Long: `Deletes a tracked issue, by default a soft delete which can be recovered, or hard delete which is permanent.`,
    Example: `    # Soft delete a specific issue by id
        issues delete 42
    # Permanently delete an issue by id
        issues delete 42 --hard`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var issueId, err = strconv.Atoi(args[0])
        if err != nil {
            fmt.Errorf("Error: invalid issue id\n")
        }

        if hardDelete {
            err := tracker.HardDelete(issueId)
            if err != nil {
            cmd.PrintErrf("Error: could not delete issue %d: %v\n", issueId, err)
                return
            }
            fmt.Printf("Permanently deleted issue %d\n", issueId)
            return
        }

        err = tracker.SoftDelete(issueId)
        if err != nil {
            cmd.PrintErrf("Error: could not delete issue %d: %v\n", issueId, err)
            return
        }
        fmt.Printf("Sucessfully deleted issue %d\n", issueId)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVar(&hardDelete, "hard", false, "Hard delete an issue (cannot be undone)")
}
