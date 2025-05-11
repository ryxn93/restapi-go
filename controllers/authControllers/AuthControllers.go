package authcontrollers

import (
	"encoding/json"
	"gin/database"
	"net/http"
	"time"

	"gin/config"
	"gin/helper"
	"gin/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(c.Writer, http.StatusBadRequest, response) // Menggunakan c.Writer untuk response
		return
	}
	defer c.Request.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau password salah"}
			helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJson(c.Writer, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(c.Writer, http.StatusInternalServerError, response)
		return
	}

	// set token di cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login berhasil cuy"}
	helper.ResponseJson(c.Writer, http.StatusOK, response)
}

func Register(c *gin.Context) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(c.Writer, http.StatusBadRequest, response)
		return
	}
	defer c.Request.Body.Close()

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := database.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(c.Writer, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJson(c.Writer, http.StatusOK, response)
}

func Logout(c *gin.Context) {
	// hapus token yang ada di cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil cuy"}
	helper.ResponseJson(c.Writer, http.StatusOK, response)
}
