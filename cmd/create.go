package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

var (
    priority string
    description string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        issueTitle := args[0]

        var chosenPriority model.Priority
        switch priority {
            case "low":
                chosenPriority = model.Low
            case "high":
                chosenPriority = model.High
            default:
                chosenPriority = model.Medium
        }

        newIssue := model.Issue{
            Title: issueTitle,
            Priority: chosenPriority,
            Description: description,
        }

        err := tracker.CreateIssue(newIssue)
        if err != nil {
            cmd.PrintErrf("Error: Failed to create issue: %v\n", err)
            return
        }

        err = tracker.UpdateGitIgnore()
        if err != nil {
            cmd.PrintErrf("Error: Failed to update .gitignore: %v\n", err)
        }

        cmd.Printf("Created new issue %s [Priority: %s]\n", newIssue.Title, newIssue.Priority)
    },
}

func init() {
    rootCmd.AddCommand(createCmd)

    createCmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Set issue importance level (low, medium, high)")
    createCmd.Flags().StringVarP(&description, "description", "d", "", "Add a detailed description of the issue")
}
