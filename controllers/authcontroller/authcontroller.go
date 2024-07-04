package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jasmineerina/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Gagal mendecode json")
	}
	defer r.Body.Close()

	// hash password menggunakan Bcrpyt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke data
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Gagal menyimpan data")
	}

	response, _ := json.Marshal(map[string]string{"message": "success"})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
