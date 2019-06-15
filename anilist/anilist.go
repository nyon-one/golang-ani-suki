package anilist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./model"
	oauth "golang.org/x/oauth2"
)

// Oauth config
func Oauth(clientID, clientSecret, redirectURL string) *oauth.Config {
	return &oauth.Config{
		RedirectURL:  redirectURL,
		ClientSecret: clientSecret,
		ClientID:     clientID,
		Endpoint: oauth.Endpoint{
			AuthURL:  "https://anilist.co/api/v2/oauth/authorize",
			TokenURL: "https://anilist.co/api/v2/oauth/token",
		},
	}
}

// API anilist
type API struct {
	Client *http.Client
}

//Req request data from api
func (c *API) Req(jsonStr string) (resp *http.Response, err error) {
	return c.Client.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer([]byte(jsonStr)))
}

//Req request data from api
func (c *API) Query(query string) (resp *http.Response, err error) {
	jsonStr := fmt.Sprintf(`{"query":"%s"}`, query)
	return c.Client.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer([]byte(jsonStr)))
}

//Get json
func (c API) Get(jsonStr string) *model.Data {
	resp, err := c.Query(jsonStr)
	if err != nil {
		log.Fatalf("\nCouldn't get data: %s\n ", err.Error())
		return nil
	}
	defer resp.Body.Close()
	var result = new(model.Data)
	json.NewDecoder(resp.Body).Decode(result)
	// if result.Errors != nil {
	// 	panic(result.Errors)
	// }
	return result
}

//New anilist api
func New(token *oauth.Token) API {
	return API{oauth.NewClient(oauth.NoContext, oauth.StaticTokenSource(token))}
}
