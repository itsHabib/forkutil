package fork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/itsHabib/art"
	"github.com/itsHabib/forkutil/api"
	"github.com/spf13/cobra"
)

type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

// ForkCmd is the command responsible for forking GitHub repositories
var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must provide a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository:", err)
		}
	},
}

// ForkRepository forks a GitHub repository
func ForkRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("Repository must be in format owner/project")
	}
	return api.GitHubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	})
}

// ForkSuccess deals with a successful forking
func ForkSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Forked to repository: %s", respContent.FullName)
	return nil
}

// GetForkResource builds a RestResource for forking
func GetForkResource() *art.RestResource {
	forkRouter := art.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(401, func(_ *http.Response) error {
		return fmt.Errorf("You must set an authentication token")
	})
	fork := art.NewResource("/repos/{{.owner}}/{{.repo}}/forks",
		"POST", forkRouter)

	return fork
}
