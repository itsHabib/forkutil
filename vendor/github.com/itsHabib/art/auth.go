package art

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// AuthToken defines a token used for auth
type AuthToken struct {
	Token string
}

// AuthBasic defines a Basic Auth structure
type AuthBasic struct {
	Username string
	Password string
}

// Authentication interface allows both implementations of basic auth
// && auth token
type Authentication interface {
	AuthorizationHeader() string
}

// NewAuthBasic returns a new AuthBasic Struct
func NewAuthBasic(username, password string) *AuthBasic {
	return &AuthBasic{
		Username: username,
		Password: password,
	}
}

// NewAuthToken returns a new AuthToken struct
func NewAuthToken(token string) *AuthToken {
	return &AuthToken{
		Token: token,
	}
}

// AuthorizationHeader returns the token header
func (at *AuthToken) AuthorizationHeader() string {
	return fmt.Sprintf("token %s", at.Token)
}

// AuthorizationHeader returns the basic auth header
func (ab *AuthBasic) AuthorizationHeader() string {
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	encContent := fmt.Sprintf("%s:%s", ab.Username, ab.Password)
	enc.Write([]byte(encContent))
	enc.Close()
	content, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Read Failed:", err)
	}
	return fmt.Sprintf("basic %s", string(content))
}
