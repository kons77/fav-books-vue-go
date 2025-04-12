package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kons77/fav-books-api/internal/data"
	"github.com/mozillazg/go-slugify"

	"github.com/go-chi/chi/v5"
)

var staticPath = "./static/" // for unite tests in the future

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

	// make sure user is active
	if user.Active == 0 {
		app.errorJSON(w, errors.New("user is not active"))
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
		Data:    envelope{"token": token, "user": user},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}

}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "logget out",
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	var users data.User
	all, err := users.GetAll()
	if err != nil {
		app.errorLog.Println()
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    envelope{"users": all},
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) EditUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if user.ID == 0 {
		// add user
		if _, err := app.models.User.InsertUser(user); err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		// edit user
		u, err := app.models.User.GetUserByID(user.ID)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		u.Email = user.Email
		u.FirstName = user.FirstName
		u.LastName = user.LastName
		u.Active = user.Active

		if err := u.UpdateUser(); err != nil {
			app.errorJSON(w, err)
			return
		}

		// if password != string , update password
		if user.Password != "" {
			err := u.ResetPassword(user.Password)
			if err != nil {
				app.errorJSON(w, err)
				return
			}
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Changes saved",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user, err := app.models.User.GetUserByID(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, user)
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.User.DeleteUserByID(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "User deleted",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) LogUserOutandSetInactive(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user, err := app.models.User.GetUserByID(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user.Active = 0
	err = user.UpdateUser()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// delete tokes for user
	err = app.models.Token.DeletTokensForUser(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "user logged out and set to inactive",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// check if token is valid
func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	valid := false
	valid, _ = app.models.Token.ValidToken(requestPayload.Token) // no need to chek error cause false when error appears

	payload := jsonResponse{
		Error: false,
		Data:  valid,
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.models.Book.GetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    envelope{"books": books},
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) OneBook(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	book, err := app.models.Book.GetOneBySlug(slug)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error: false,
		Data:  book,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AuthorsAll(w http.ResponseWriter, r *http.Request) {
	all, err := app.models.Author.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	type selectData struct {
		Value int    `json:"value"`
		Text  string `json:"text"`
	}

	var results []selectData

	for _, x := range all {
		author := selectData{
			Value: x.ID,
			Text:  x.AuthorName,
		}

		results = append(results, author)
	}

	payload := jsonResponse{
		Error: false,
		Data:  results,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) EditBook(w http.ResponseWriter, r *http.Request) {
	// can't use type Book cause of more fields
	var requestPayload struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		AuthorID        int    `json:"author_id"`
		PublicationYear int    `json:"publication_year"`
		Description     string `json:"description"`
		CoverBase64     string `json:"cover"`
		GenreIDs        []int  `json:"genre_ids"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book := data.Book{
		ID:              requestPayload.ID,
		Title:           requestPayload.Title,
		AuthorID:        requestPayload.AuthorID,
		PublicationYear: requestPayload.PublicationYear,
		Description:     requestPayload.Description,
		Slug:            slugify.Slugify(requestPayload.Title),
		GenreIDs:        requestPayload.GenreIDs,
	}

	if len(requestPayload.CoverBase64) > 0 {
		// we have a cover
		decoved, err := base64.StdEncoding.DecodeString(requestPayload.CoverBase64)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		// write image to /static/covers
		if err := os.WriteFile(fmt.Sprintf("%s/covers/%s.jpg", staticPath, book.Slug), decoved, 0666); err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	if book.ID == 0 {
		// adding a book
		_, err := app.models.Book.Insert(book)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		// updating a book
		err := book.Update()
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Changes Saved",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) BookByID(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book, err := app.models.Book.GetOneById(bookID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    book,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.Book.DeleteByID(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Book deleted",
	}

	app.writeJSON(w, http.StatusOK, payload)
}
