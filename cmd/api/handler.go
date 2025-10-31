package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dmc0001/auth-jwt-project/internal/types"
	"github.com/dmc0001/auth-jwt-project/internal/utils"
)

func (app *Application) getUser(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("param")

	if strings.Contains(param, "@") {
		// treat as email
		app.getUserByEmail(w, param)
		return
	}

	// else treat as ID
	id, err := strconv.Atoi(param)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}

	app.getUserById(w, id)

}

func (app *Application) getUserByEmail(w http.ResponseWriter, email string) {

	user, err := app.config.userModel.GetUserByEmail(email)
	if err != nil {
		utils.WritingError(w, http.StatusNotFound, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, user)
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}
func (app *Application) getUserById(w http.ResponseWriter, id int) {

	user, err := app.config.userModel.GetUserById(id)
	if err != nil {
		utils.WritingError(w, http.StatusNotFound, err)
		return
	}

	err = utils.ParsingToJson(w, http.StatusOK, user)
	if err != nil {
		utils.WritingError(w, http.StatusInternalServerError, err)
	}

}

func (app *Application) loginUser(w http.ResponseWriter, r *http.Request) {
	// get user payload
	var loginCredentials types.LoginUserRequest
	err := utils.ParsingFromJson(r, &loginCredentials)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
	}
	user, err := app.config.userModel.LoginUser(loginCredentials)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}
	utils.ParsingToJson(w, http.StatusOK, &user)
}

func (app *Application) registerUser(w http.ResponseWriter, r *http.Request) {
	// get user payload
	var registerCredentials types.RegisterUserRequest
	err := utils.ParsingFromJson(r, &registerCredentials)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}
	// check if user exist and if does exist create new one
	err = app.config.userModel.RegisterUser(registerCredentials)
	if err != nil {
		utils.WritingError(w, http.StatusBadRequest, err)
		return
	}
	utils.ParsingToJson(w, http.StatusCreated, map[string]string{
		"message": "✅ User registered successfully",
	})

}
