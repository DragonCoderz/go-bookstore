package utils

//Function that facilitates the transfer of data via jsons, http requests, etc.

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}