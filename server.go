package main

import (
	"fmt"
	"net/http"

	//third party packages
	"github.com/julienschmidt/httprouter"
)

func main() {
	//instantiate a new router
	r := httprouter.New()

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome\n")
	})

	http.ListenAndServe("localhost:3000", r)

}
