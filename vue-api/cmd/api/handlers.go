package main

import (
	"errors"
	"net/http"
	"time"
)

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type envelope map[string]any

// Login is the handler used to attempt to log a user into the api
func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"` //email
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {

		app.errorLog.Println()
		//app.errorJSON(w, err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// TODO authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// look up user by email
	user, err := app.models.User.GetUserByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// validate the user's password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// we have a valid use, generate a token
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// save it to db
	err = app.models.Token.InsertToken(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//send back a response
	payload = jsonResponse{
		Error:   false,
		Message: "logged in in",
		Data:    envelope{"token": token},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}

}

/*
// AdminPostHashPassword generate hashed password from the string
func (app *application) AdminPostHashPassword(w http.ResponseWriter, r *http.Request) {


	form := forms.New(r.PostForm)
	form.Required("password")

	if !form.Valid() {
		// send json response
		resp := map[string]string{
			"error": "Password cannot be empty!",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	pswd := r.Form.Get("password")

	hashedPswd, err := helpers.HashPassword(pswd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	resp := map[string]string{
		"HashedPswd": string(hashedPswd),
	}

	// send json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) // writes directly to http.ResponseWriter: No need for intermediate variables.

}

*/
