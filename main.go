package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	// token := os.Getenv("GITHUB_TOKEN")
	// maxAssigned := 5
	// maxOpenPRs := 3
	// readyLabel := "contributor-assign-ready"
	// assignedLabel := "assigned"

	// Create GitHub client
	// client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))

	payload := os.Getenv("GITHUB_EVENT_PATH")

	event, err := github.ParseWebHook("issue_comment", []byte(payload))
	if err != nil {
		fmt.Println("Error parsing webhook:", err)
		return
	}

	// Check if the event is an issue comment event
	if commentEvent, ok := event.(*github.IssueCommentEvent); ok {
		// Get the username of the commenter
		username := commentEvent.GetComment().GetUser().GetLogin()
		fmt.Println("Username:", username)
	}

}
