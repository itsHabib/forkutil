package art

import "fmt"

// API is used to map strings to specific REST resources
// Simplifies making requests by using a simple string name
// to make requests
type API struct {
	BaseURL       string
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

// NewAPI instantiates a new API struct given a baseURL
func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}

// SetAuth sets authentication on API client
func (api *API) SetAuth(auth Authentication) {
	api.Client.AuthInfo = auth
}

// AddResource adds resources to resources map
func (api *API) AddResource(name string, res *RestResource) {
	api.Resources[name] = res
}

// Call uses client to make a http request on a given resource
// given a name
func (api *API) Call(name string, params map[string]string) error {
	res, ok := api.Resources[name]
	if !ok {
		return fmt.Errorf("Resource does not exist: %s", name)
	}
	return api.Client.ProcessRequest(api.BaseURL, res, params)
}

// ResourceNames returns resources in API struct
func (api *API) ResourceNames() []string {
	resources := []string{}
	for k := range api.Resources {
		resources = append(resources, k)
	}
	return resources
}
