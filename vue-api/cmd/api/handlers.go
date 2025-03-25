package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"username"`
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

	//send back a response
	payload.Error = false
	payload.Message = "Signed in"

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
