package controllers

import (
	"encoding/json"
	// "golang/configs"
	// "golang/helpers"
	"golang/configs"
	"golang/helpers"
	"golang/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Password tidak sama", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Username: register.Username,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Register berhasil", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var login models.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "username = ?", login.Username).Error; err != nil {
		helpers.Response(w, 404, "Username atau password salah", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Username atau password salah", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Login berhasil", token)
}
