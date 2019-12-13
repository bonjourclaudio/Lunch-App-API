package http

import (
	"github.com/claudioontheweb/lunch-app/config"
	"github.com/claudioontheweb/lunch-app/pkg/api/auth"
	"github.com/claudioontheweb/lunch-app/pkg/api/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Services struct {
	userService services.UserService
	authService auth.AuthService
}

type Server struct {
	Services 	Services
	Router		*mux.Router
	Config		*config.ServerConfig
}

func NewServer(userService services.UserService, authService auth.AuthService, config *config.ServerConfig) *Server {

	serv := Services{userService:userService,authService: authService}

	server := Server{
		Services: serv,
		Router:   nil,
		Config:   config,
	}

	return &server
}

func (api *Server) RegisterHandlers() {

	r := mux.NewRouter()

	r.Use(api.corsMiddleware)
	r.Use(api.commonMiddleware)
	r.Use(api.loggingMiddleware)

	api.NewUserHandler(r.PathPrefix("/users").Subrouter())
	api.NewAuthHandler(r)

	api.Router = r

}

func (api *Server) StartHTTP() error {

	server := &http.Server {
		Addr:              api.Config.Host + ":" + api.Config.Port,
		Handler:           api.Router,
	}

	return server.ListenAndServe()

}


func (api *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		next.ServeHTTP(w, r)

	})
}

func (api *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.WithFields(log.Fields{
			"request_uri": r.RequestURI,
			"protocol": r.Proto,
			"method": r.Method,
		}).Debug("Calling Endpoint")

		next.ServeHTTP(w, r)

	})
}

func (api *Server) commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)

	})
}

