package foursquare

// https://developer.foursquare.com/docs/responses/venue

type Venue struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Location   Location   `json:"location"`
	Categories []Category `json:"categories"`
}

type Location struct {
	Address     string  `json:"address"`
	CrossStreet string  `json:"crossStreet"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Distance    float64 `json:"distance"`
}

type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Primary    bool   `json:"primary"`
}
