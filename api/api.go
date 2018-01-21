package api

import (
	"github.com/itsHabib/art"
	"github.com/itsHabib/forkutil/fork"
	"github.com/spf13/viper"
)

var api *art.API

// GitHubAPI returns
func GitHubAPI() *art.API {
	if api == nil {
		api = art.NewAPI("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(art.NewAuthToken(token))
		api.AddResource("fork", fork.GetForkResource())
	}
	return api
}
