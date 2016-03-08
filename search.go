package foursquare

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SearchReq struct {
	LatLng string // Latitude and longitude of the user's location.
	Near   string // A string naming a place in the world.
	Query  string // A search term to be applied against venue names.
	Limit  string // Number of results to return, up to 50.
	Radius int    // The maximum supported radius is currently 100,000 meters.

	// One of the values below, indicating your intent in performing the search. If no value is specified, defaults to checkin.
	//    checkin	Finds results that the current user (or, for userless requests, a typical user) is likely to check in to at the provided ll at the current moment in time. This is the intent we recommend most apps use.
	//    browse	Find venues within a given area. Unlike the checkin intent, browse searches an entire region instead of only finding Venues closest to a point. You must define a region to search be including either the ll and radius parameters, or the sw and ne. The region will be a spherical cap if you include the ll and radius parameters, or it will be a bounding quadrangle if you include the sw and ne parameters.
	//    global	Finds the most globally relevant venues for the search, independent of location. Ignores all other parameters other than query and limit.
	//    match	Finds venues that are are nearly-exact matches for the given parameters. This intent is highly sensitive to the provided location. We recommend using this intent only when trying to correlate an existing place database with Foursquare's. The results will be sorted best match first, taking distance and spelling mistakes/variations into account.
	//          query and ll are the only required parameters for this intent, but matching also supports phone, address, city, state, zip, and twitter. There's no specified format for these parameters—we do our best to normalize them and drop them from the search if unsuccessful.
	Intent string

	// With ne, limits results to the bounding quadrangle defined by the latitude and longitude given by sw as its south-west corner, and ne as its north-east corner. The bounding quadrangle is only supported for intent=browse searches. Not valid with ll or radius. Bounding quadrangles with an area up to approximately 10,000 square kilometers are supported.
	SW string
	NE string
}

type Meta struct {
	Code       int    `json:"code"`
	ErrorType  string `json:"errorType"`
	ErorDetail string `json:"errorDetail"`
}

type FSResp struct {
	Meta Meta       `json:"meta"`
	Resp SearchResp `json:"response"`
}

type SearchResp struct {
	Venue []Venue `json:"venues"`
}

func (api *Api) Search(req *SearchReq) (*FSResp, error) {

	// set param
	params := api.clientParams()

	params.Set("v", "20140806")

	if len(req.LatLng) > 0 {
		params.Set("ll", req.LatLng)
	} else {
		params.Set("near", req.Near)
	}

	url := SEARCHURL + params.Encode()
	log.Println(url)

	http.DefaultClient.Timeout = 5 * time.Second // 超时 5s
	r, e := http.Get(url)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer r.Body.Close()

	// result
	resp := &FSResp{}
	dec := json.NewDecoder(r.Body)
	if e := dec.Decode(resp); e != nil {
		log.Println(e)
		return nil, e
	}

	return resp, nil
}
