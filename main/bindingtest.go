package main

import (
	"net/http"
	"tripod/net"
	"Work/main/bind"
)

type bindQueryField struct {
	IntField string `form:"int_field"`
	BoolField bool `form:"bool_field"`
	stringField bool `form:"string_field"`
	sliceField []string `form:"array_field"`
}

func main()  {
	mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	reqQueryFields := bindQueryField{}
    	err := bind.ShouldBind(r, reqQueryFields)
    	if err != nil {
			net.WriteError(w,0,"数据校验错误")
		}
		net.WriteJSON(w,reqQueryFields)
	})
	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}
	server.ListenAndServe()
}
