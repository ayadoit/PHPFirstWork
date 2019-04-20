package bind

import (
	"net/http"
)

const defaultMemory = 32 * 1024 * 1024

type formBinding struct{}

func (formBinding) ContentType() string {
	return "form"
}

func (formBinding) Bind(r *http.Request, obj interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	r.ParseMultipartForm(defaultMemory)

	return mapForm(obj, r.Form)
}
