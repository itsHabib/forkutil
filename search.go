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

// SearchResponse holds the response after a search request to GH
// is made
type SearchResponse struct {
	Results []*SearchResult `json:"items"`
}

// SearchResult is used to hold each search result
type SearchResult struct {
	FullName string `json:"full_name"`
}

// SearchCmd represents the command used to search GitHub repositories
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword(args); err != nil {
			log.Fatalln("Search Failed:", err)
		}
	},
}

// SearchByKeyword searches GitHub repositories by keyword
func SearchByKeyword(keywords []string) error {
	return GitHubAPI().Call("search", map[string]string{
		"query": strings.Join(keywords, "+"),
	})
}

// SearchSuccess is the function to deal with successful searches from GH
func SearchSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := SearchResponse{}
	json.Unmarshal(content, &respContent)
	for _, item := range respContent.Results {
		fmt.Println(item.FullName)
	}
	return nil
}

// SearchDefaultRouter returns an error
func SearchDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

// GetSearchResource is used to create a search resource
func GetSearchResource() *art.RestResource {
	searchRouter := art.NewRouter()
	searchRouter.DefaultRouter = SearchDefaultRouter
	searchRouter.RegisterFunc(200, SearchSuccess)
	search := art.NewResource("/search/repositories?q={{.query}}",
		"GET", searchRouter)
	return search
}
