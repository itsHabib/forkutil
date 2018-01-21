package art

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

// RestResource is the struct containing the REST endpoints
// corresponding methods, and corresponding routers for that resource
type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}

// NewResource instantiates a new REST resource
func NewResource(endpoint, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

// RenderEndpoint parses params and constructs an endpoint
func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}
	t, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalln("unalbe to parse endpoint")
	}
	buffer := &bytes.Buffer{}
	t.Execute(buffer, params)
	endpoint, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Unable to read endpoint")
	}
	return string(endpoint)
}
