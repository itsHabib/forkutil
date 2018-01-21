package art

import (
	"net/http"
	"strings"
)

// Client holds the http client and either an auth basic or token
// implementation
type Client struct {
	Client   *http.Client
	AuthInfo Authentication
}

// NewClient instantiates a new client and leaves no authentication
// by default in case it is not neededs
func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

// SetAuth sets the type of authentication for the Client struct
func (c *Client) SetAuth(auth Authentication) {
	c.AuthInfo = auth
}

// ProcessRequest uses client to send a request and call an callback router
func (c *Client) ProcessRequest(baseURL string, res *RestResource,
	params map[string]string) error {
	endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	url := trimmedBaseURL + "/" + endpoint
	req, err := http.NewRequest(res.Method, url, nil)
	if err != nil {
		return err
	}
	if c.AuthInfo != nil {
		req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	return res.Router.CallFunc(resp, nil)
}
