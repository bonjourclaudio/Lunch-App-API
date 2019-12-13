package http

import (
	"encoding/json"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (api *Server) NewUserHandler(r *mux.Router) {
	r.HandleFunc("", api.FindAll).Methods("GET")
	r.HandleFunc("/:id", api.Find).Methods("GET")
	r.HandleFunc("/:id", api.Update).Methods("UPDATE")
	r.HandleFunc("/:id", api.Delete).Methods("DELETE")
	r.Use(api.authMiddleware)
}

func (api *Server) FindAll(w http.ResponseWriter, r *http.Request) {

	users := api.Services.userService.GetAllUsers()

	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (api *Server) Find(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	user := api.Services.userService.GetUserById(id)

	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (api *Server) Update(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(api.Services.userService.UpdateUser(user))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (api *Server) Delete(w http.ResponseWriter, r *http.Request) {


}