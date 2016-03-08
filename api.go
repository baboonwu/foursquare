package foursquare

import (
	"net/url"
)

const (
	EXPLOREURL = "https://api.foursquare.com/v2/venues/explore?"
	SEARCHURL  = "https://api.foursquare.com/v2/venues/search?"
)

type Api struct {
	ClientId     string
	ClientSecret string
}

func NewApi(clientId string, clientSecret string) *Api {
	return &Api{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}
func (api *Api) clientParams() url.Values {
	return url.Values{
		"client_id":     []string{api.ClientId},
		"client_secret": []string{api.ClientSecret},
	}
}
