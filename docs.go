package forkutil

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/itsHabib/art"
	"github.com/spf13/cobra"
)

// ReadmeResponse holds the contect of readmes
type ReadmeResponse struct {
	Content string `json:"content"`
}

// DocsCmd represents the command used to get docs from a GitHub repository
var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Read the documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("Must supply repository argument")
		}
		if err := GetRepositoryReadme(args[0]); err != nil {
			log.Fatalln("Failed to get docs: ", err)
		}
	},
}

// GetRepositoryReadme returns a repositories readme
// TODO implement actual function using GitHub API
func GetRepositoryReadme(repository string) error {
	values := strings.Split(repository, "/")
	return GitHubAPI().Call("docs", map[string]string{
		"owner":   values[0],
		"project": values[1],
	})
}

func ReadmeSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := ReadResponse{}
	json.Unmarshal(content, &respContent)
	buff, err := base64.StdEncoding.DecodeString(respContent.Content)
	if err != nil {
		return err
	}
	fmt.Println(string(buff))
	return nil
}

// ReadmeDefaultRouter is the default router for readmes
func ReademeDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

// GetReadmeResource creates a Rest resource for getting readmes from repos
func GetReadmeResource() *art.RestResource {
	router := art.NewRouter()
	router.RegisterFunc(200, ReadmeSuccess)
	router.DefaultRouter = ReadmeDefaultRouter
	readmeResource := art.NewResource("/repos/{{.owner}}/{{.project}}/readme",
		"GET", router)
	return readmeResource
}
