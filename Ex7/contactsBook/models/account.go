package models

import (
	u "contactsBook/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if resp, ok := account.ValidateData(); !ok {
		return resp, false
	}

	acc := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error!, Error: "+err.Error()), false
	}
	if acc.Email != "" {
		return u.Message(false, "The email is already occupied by another user!"), false
	}

	return u.Message(false, "Check is passed!"), true
}

func (account *Account) ValidateData() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is not valid!"), false
	}

	if len(account.Password) < 4 {
		return u.Message(false, "Password must be longer then 4 symbols"), false
	}

	return u.Message(false, "Check is passed!"), true
}

func (account *Account) CreateAccount() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	account.Password = ""

	response := u.Message(true, "Account has been created!")
	response["account"] = account
	return response
}

func LoginAccount(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password does not match!")
	}

	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func UpdateAccount(id uint, email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("id = ?", id).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Account id not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}
	account.Email = email
	account.Password = password

	if resp, ok := account.ValidateData(); !ok {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Save(account)

	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(true, "Account updated")
	resp["account"] = account
	return resp
}

func DeleteAccount(id uint) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("id = ?", id).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Account id not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	GetDB().Delete(account).Where("id = ", id)

	resp := u.Message(true, "Account deleted")
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
