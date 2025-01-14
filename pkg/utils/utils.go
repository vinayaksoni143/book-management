package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Parses the body of the request and returns the body as a map
func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}
