package http

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/claudioontheweb/lunch-app/pkg/api/auth"
	"github.com/claudioontheweb/lunch-app/pkg/db/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (api *Server) NewAuthHandler(r *mux.Router) {
	r.HandleFunc("/login", api.Login).Methods("GET")
	r.HandleFunc("/callback", api.Callback).Methods("GET")
}

func (api *Server) Login(w http.ResponseWriter, r *http.Request) {

	oauthStateString := api.generateOAuthStateCookie(w)

	url := api.Services.authService.Login().Config.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (api *Server) generateOAuthStateCookie(w http.ResponseWriter) string {

	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state

}


func (api *Server) Callback(w http.ResponseWriter, r *http.Request) {

	oauthStateString, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthStateString.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user, err := getUserInfo(r.FormValue("code"), api)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Check if user is already registered
	userExists, existingUser := api.Services.userService.CheckUser(user.Email)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate token
	token, err := api.generateToken(user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authRes := auth.AuthResponse{}
	authRes.Token = token

	if userExists {
		// Return existing user
		authRes.User = *existingUser
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&authRes)
		return

	} else {
		// Create new user in DB
		newUser := api.Services.userService.CreateUser(user)
		authRes.User = *newUser
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&authRes)
		return
	}
}

func getUserInfo(code string, api *Server) (*models.User, error) {

	token, err := api.Services.authService.Login().Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	var user *models.User
	_ = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return user, nil

}

func (api *Server)generateToken(user *models.User) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(api.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (api *Server) authMiddleware(next http.Handler) http.Handler {

		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
				return []byte(api.Config.JWTSecret), nil
			},
			SigningMethod:       jwt.SigningMethodHS256,
		})

		return jwtMiddleware.Handler(next)
}