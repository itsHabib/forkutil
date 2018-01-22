package forkutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/itsHabib/art"
	"github.com/spf13/cobra"
)

// PullRequestPayload defines the structure needed to send a pull request
type PullRequestPayload struct {
	Title        string `json:"title"`
	Message      string `json:"body"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       bool   `json:"maintainer_can_modigy"`
}

// PullRequestResponse holds the url for the pull request
type PullRequestResponse struct {
	URL string `json:"html_url"`
}

// PullRequestCmd defines the command line command to execute pull requests
var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create a Pull Request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreatePullRequest(); err != nil {
			log.Fatalln("Failed to create pull request:", err)
		}
	},
}

func CreatePullRequest() error {
	sourceValues := strings.Split(sourceRepo, ":")
	if !(len(sourceValues) == 1 || len(sourceValues) == 2) {
		return fmt.Errorf("Source repository must be in the format [owner: ]branch got %v", sourceRepo)
	}
	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		return fmt.Errorf("Destination repository must be in the format owner/project:branch got %v",
			destRepo)
	}
	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		return fmt.Errorf("Destination repository must be in format owner/prject:branch got %v",
			destValues)
	}
	payload := &PullRequestPayload{
		Title:        pullRequestTitle,
		Message:      pullRequestMessage,
		SourceBranch: sourceRepo,
		Modify:       true,
		DestBranch:   destBranchValues[1],
	}
	return GitHubAPI().Call("pullrequest", map[string]string{
		"owner":   destValues[0],
		"project": destValues[1],
	}, payload)
}

func PullRequestSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := PullRequestResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Created Pull Request: %s", respContent.URL)
	return nil
}

// PullRequestDefaultRouter is the default router for pull request commands
func PullRequestDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("Status code %d", resp.StatusCode)
}

// GetPullRequestResource returns a rest rource to create pull requests
func GetPullRequestResource() *art.Resource {
	router := art.NewRouter()
	router.RegisterFunc(201, PullRequestSuccess)
	router.DefaultRouter = PullRequestRouter
	resource := art.NewResource("/repos/{{.owner}}/{{.project}}/pulls",
		"POST", router)
	return resource
}

// sourceRepo and destRepo hold values passed to the pull request command falgs
var sourceRepo string
var destRepo string
var pullRequestTitle string
var pullRequestMessage string

func init() {
	PullRequestCmd.Flags().StringVarP(&sourceRepo, "source", "s", "",
		"source repository")
	PullRequestCmd.Flags().StringVarP(&destRepo, "destination", "d", "",
		"destination repository")
	PullRequestCmd.Flags().StringVarP(&pullRequestTitle, "title", "t", "Basic Pull Request",
		"pull request title")
	PullRequestCmd.Flags().StringVarP(&pullRequestMessage, "message", "m", "Pull Request Message",
		"pull request message")

}
