package cmd

import (
    "time"
    "strings"

	"github.com/spf13/cobra"
    "issue-tracker/model"
    "issue-tracker/tracker"
)

var (
    createPriority string
    createDescription string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new issue",
	Long: `Create and save a new issue. By default, issue priority will be set to [medium]`,
    Example: `    # Create an issue with title 'fix bug' and high priority
        issues create "fix bug" --priority high
    # Create an issue with title 'add feature' and description 'add the xyz feature'
        issues create "add feature" --description "add the xyz feature"`,
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        issueTitle := args[0]
        
        createPriority = strings.ToLower(createPriority)

        if !IsValidPriority(createPriority) {
            cmd.PrintErrf("Error: Invalid priority %q. Must be one of: %v\n", createPriority, GetValidPriorities())
            return
        }

        chosenPriority := ParsePriority(createPriority)
        
        currentTime := time.Now()

        newIssue := model.Issue{
            Title: issueTitle,
            Priority: chosenPriority,
            Description: createDescription,
            CreatedAt: currentTime,
            LastUpdated: currentTime,
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

        cmd.Printf("Created new issue \"%s\" [Priority: %s]\n", newIssue.Title, newIssue.Priority)
    },
}

func init() {
    rootCmd.AddCommand(createCmd)

    createCmd.Flags().StringVarP(&createPriority, "priority", "p", "medium", "Set issue importance level (low, medium, high)")
    createCmd.Flags().StringVarP(&createDescription, "description", "d", "", "Add a description of the issue")
}
