package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TimCoding/rest-go-api-tutorial/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Bob Smith",
		Gender: "male",
		Age:    50,
		Id:     p.ByName("id"),
	}

	//Marshal to put user info into json format

	uj, _ := json.Marshal(u)

	//Write constant-type, status code, payload

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	//Add an ID

	u.Id = "foo"

	//Marshal newly created user struct into json

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}
