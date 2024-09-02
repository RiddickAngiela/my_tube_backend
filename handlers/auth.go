package handlers

import (
    "encoding/json"
    "net/http"
    "my_tube_backend/models"
    "my_tube_backend/utils"
    "golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    // Save user to the database (pseudo-code)
    // db.Create(&user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Fetch user from the database (pseudo-code)
    // var dbUser models.User
    // db.Where("email = ?", user.Email).First(&dbUser)

    // if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
    //     http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    //     return
    // }

    token, err := utils.GenerateJWT(user.Email)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Authorization", "Bearer "+token)
    w.WriteHeader(http.StatusOK)
}
