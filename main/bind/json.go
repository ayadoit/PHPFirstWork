package bind

import (
	"encoding/json"
	"net/http"
)

type jsonBinding struct{}

func (jsonBinding) ContentType() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(req.Body)

	return decoder.Decode(obj)
}
