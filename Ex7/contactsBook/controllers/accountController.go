package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}
	resp := account.CreateAccount()
	u.Respond(w, resp)
}

var LoginAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}
	resp := models.LoginAccount(account.Email, account.Password)
	u.Respond(w, resp)
}

var UpdateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}
	id := r.Context().Value("user").(uint)

	resp := models.UpdateAccount(id, account.Email, account.Password)
	u.Respond(w, resp)
}

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	resp := models.DeleteAccount(id)
	u.Respond(w, resp)
}
