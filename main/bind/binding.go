package bind

import (
	"net/http"
)

type Binding interface {
	ContentType() string
	Bind(*http.Request, interface{}) error
}

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

var (
	JSON = jsonBinding{}
	Form = formBinding{}
)

func GetBindingByContentType(method, contentType string) Binding {
	if method == "GET" {
		return Form
	}

	switch contentType {
	case MIMEJSON:
		return JSON
	default: //case MIMEPOSTForm, MIMEMultipartPOSTForm:
		return Form
	}
}

func ShouldBind(r *http.Request, obj interface{}) error {
	contentType := ""
	if len(r.Header["Content-Type"]) > 0 {
		contentType = r.Header["Content-Type"][0]
	}

	binding := GetBindingByContentType(r.Method, contentType)

	return binding.Bind(r, obj)
}
