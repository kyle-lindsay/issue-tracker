package cmd

import (
    "strconv"
    "strings"

	"github.com/spf13/cobra"
    "issue-tracker/tracker"
)

var recoverStatus string 

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover the deleted issue with specified id",
	Long: `Recovers a deleted issue with specified id. Sets status to [Open] by default. Note, cannot recover hard deleted issues.`,
    Example: `    # Recover a specific deleted issue by id, setting status to [Open]
        issues recover 42
    # Recover a specific issue by id, setting status to [Closed]
        issues recover 42 --status closed`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        var issueId, err = strconv.Atoi(args[0])
        if err != nil {
            cmd.PrintErrf("Error: invalid issue id\n")
            return
        }

        
        recoverStatus = strings.ToLower(recoverStatus)
        if !IsValidStatus(recoverStatus) {
            cmd.PrintErrf("Error: invalid status: %q. Must be one of: %v\n", recoverStatus, GetValidStatuses())
            return
        }

        chosenStatus := ParseStatus(recoverStatus)

        err = tracker.RecoverSoftDelete(issueId, chosenStatus)
        if err != nil {
            cmd.PrintErrf("Error: failed to recover issue: %v\n", err)
            return
        }

        cmd.Printf("Recovered issue #%d\n", issueId)
    },
}

func init() {
	rootCmd.AddCommand(recoverCmd)

    recoverCmd.Flags().StringVarP(&recoverStatus, "status", "s", "open", "Specify a status to set the recovered issue to. Default 'Open'")
}
