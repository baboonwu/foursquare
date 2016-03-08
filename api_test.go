package foursquare

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const CLIENTID = ""
const CLIENTSECRET = ""

func Test_SearchNear(t *testing.T) {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	api := NewApi(CLIENTID, CLIENTSECRET)
	resp, e := api.Search(&SearchReq{Near: "朝阳区"})
	if e != nil {
		t.Error("Test_SearchNear")
		return
	}
	fmt.Println(resp.Meta)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
}

func Test_SearchLatLng(t *testing.T) {
	api := NewApi(CLIENTID, CLIENTSECRET)
	resp, e := api.Search(&SearchReq{LatLng: "31.33,109.312"})
	if e != nil {
		t.Error("Test_SearchLatLng")
		return
	}
	fmt.Println(resp.Meta)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
}
