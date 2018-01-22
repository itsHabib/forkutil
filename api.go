package forkutil

import (
	"github.com/spf13/viper"
	"github.com/itsHabib/art"
)

var api *art.API

// GitHubAPI returns
func GitHubAPI() *art.API {
	if api == nil {
		api = art.NewAPI("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(art.NewAuthToken(token))
		api.AddResource("fork", GetForkResource())
		api.AddResource("search", GetSearchResource())
		api.AddResource("docs", GetReadmeResource())
		api.AddResource("pullrequest", GetPullRequestResource())
	}
	return api
}
