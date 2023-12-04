package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data, resp := models.GetContacts(id)
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		ContactId uint           `json:"id"`
		Contact   models.Contact `json:"contact"`
	}
	data := &Body{}
	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}

	resp := models.UpdateContatct(data.ContactId, data.Contact)
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		ID uint `json:"id"`
	}
	data := &Body{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		u.Respond(w, u.ErrorMessage(false, "Bad request", 400))
		return
	}
	resp := models.DeleteContact(data.ID)
	u.Respond(w, resp)
}
