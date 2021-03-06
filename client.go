package tesla

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Required authorization credentials for the Tesla API
type Auth struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

// The token and related elements returned after a successful auth
// by the Tesla API
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Expires     int64
}

// Provides the client and associated elements for interacting with the
// Tesla API
type Client struct {
	Auth         *Auth
	Token        *Token
	HTTP         *http.Client
	BaseURL      string
	StreamingURL string
}

var AuthURL = "https://owner-api.teslamotors.com/oauth/token"

const BaseURL = "https://owner-api.teslamotors.com/api/1"

// Generates a new client for the Tesla API
func NewClient(auth *Auth) (*Client, error) {
	client := &Client{
		Auth:         auth,
		HTTP:         &http.Client{},
		BaseURL:      BaseURL,
		StreamingURL: StreamingURL,
	}
	token, err := client.authorize(auth)
	if err != nil {
		return nil, err
	}
	client.Token = token
	return client, nil
}

// NewClientWithToken Generates a new client for the Tesla API using an existing token
func NewClientWithToken(auth *Auth, token *Token) (*Client, error) {
	client := &Client{
		Auth:         auth,
		HTTP:         &http.Client{},
		Token:        token,
		BaseURL:      BaseURL,
		StreamingURL: StreamingURL,
	}
	if client.TokenExpired() {
		return nil, errors.New("supplied token is expired")
	}
	return client, nil
}

// TokenExpired indicates whether an existing token is within an hour of expiration
func (c Client) TokenExpired() bool {
	exp := time.Unix(c.Token.Expires, 0)
	return time.Until(exp) < time.Duration(1*time.Hour)
}

// Authorizes against the Tesla API with the appropriate credentials
func (c Client) authorize(auth *Auth) (*Token, error) {
	now := time.Now()
	auth.GrantType = "password"
	data, _ := json.Marshal(auth)
	body, err := c.post(AuthURL, data)
	if err != nil {
		return nil, err
	}
	token := &Token{}
	if err := json.Unmarshal(body, token); err != nil {
		return nil, err
	}
	token.Expires = now.Add(time.Second * time.Duration(token.ExpiresIn)).Unix()
	return token, nil
}

// Calls an HTTP GET
func (c Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	return c.processRequest(req)
}

// getJSON performs an HTTP GET and then unmarshals the result into the provided struct.
func (c Client) getJSON(url string, out interface{}) error {
	body, err := c.get(url)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, out); err != nil {
		return err
	}
	return nil
}

// Calls an HTTP POST with a JSON body
func (c Client) post(url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// Processes a HTTP POST/PUT request
func (c Client) processRequest(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Sets the required headers for calls to the Tesla API
func (c Client) setHeaders(req *http.Request) {
	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
