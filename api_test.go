package foursquare

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const CLIENTID = ""
const CLIENTSECRET = ""

func Test_SearchNear(t *testing.T) {
	api := NewApi(CLIENTID, CLIENTSECRET)
	resp, _ := api.Search(&SearchReq{Near: "朝阳区"})
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
	fmt.Println(resp.Meta)
}

func Test_SearchLatLng(t *testing.T) {
	api := NewApi(CLIENTID, CLIENTSECRET)
	resp, _ := api.Search(&SearchReq{LatLng: "31.33,109.312"})
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
	fmt.Println(resp.Meta)
}
