package main

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}
	user, err := app.DB.GetUserByEmail(creds.Username)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

}
func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {

}
func (app *application) getUser(w http.ResponseWriter, r *http.Request) {

}
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

}
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {

}
func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {

}
