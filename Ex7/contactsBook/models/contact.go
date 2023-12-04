package models

import (
	u "contactsBook/utils"
	"fmt"
	"unicode"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty!"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number cannot be empty!"), false
	}

	for _, v := range contact.Name {
		if unicode.IsDigit(v) {
			return u.Message(false, "Invalid name!"), false
		}
	}

	for _, v := range contact.Phone {
		if !unicode.IsDigit(v) {
			return u.Message(false, "Invalid phone number!"), false
		}
	}

	if len(contact.Phone) != 11 {
		return u.Message(false, "Phone number cannot be empty!"), false
	}

	return u.Message(true, "Check is passed!"), true
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Create(contact)

	resp := u.Message(true, "Contact created")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}

func UpdateContatct(id uint, data Contact) map[string]interface{} {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	contact.Phone = data.Phone
	contact.Name = data.Name
	contact.UserId = data.UserId

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Save(contact)

	resp := u.Message(true, "Contact updated")
	resp["contact"] = contact
	return resp
}

func DeleteContact(id uint) map[string]interface{} {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	GetDB().Delete(contact).Where("id = ?", id)

	resp := u.Message(true, "Contact deleted")
	return resp
}
