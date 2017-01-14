package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//third party packages
	"github.com/TimCoding/rest-go-api-tutorial/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//instantiate a new router
	r := httprouter.New()

	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		u := models.User{
			Name:   "Bob Smith",
			Gender: "male",
			Age:    50,
			Id:     p.ByName("id"),
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s\n", uj)

	})

	r.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//Empty user to be populated
		u := models.User{}

		//Populate from body
		json.NewDecoder(r.Body).Decode(&u)

		//Add id
		u.Id = "foo"

		//Marshal provided interface into json

		uj, _ := json.Marshal(u)

		//Write content-type, status code, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s\n", uj)
	})

	r.DELETE("/user/:id", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(200)
	})

	http.ListenAndServe("localhost:3000", r)

}
