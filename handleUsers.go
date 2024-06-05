package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/Appointment-App/auth"
	"github.com/stonoy/Appointment-App/internal/database"
)

func (cfg *apiConfig) register(w http.ResponseWriter, r *http.Request) {
	// get the request data
	type reqStruct struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// check all inputs
	if reqObj.Name == "" || reqObj.Email == "" || reqObj.Password == "" || len(reqObj.Password) < 6 {
		respWithError(w, 400, "Follow the input instructions")
		return
	}

	// hash the password
	hashedPassword, err := auth.GenerateHashedPassword(reqObj.Password)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error hashing password : %v", err))
		return
	}

	// check for admin
	isUserSetForAdmin, err := cfg.DB.IsSetForAdmin(r.Context())
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in IsSetForAdmin : %v", err))
		return
	}

	userRole := "patient"
	if isUserSetForAdmin {
		userRole = "admin"
	}

	// create new user
	theNewUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      reqObj.Name,
		Email:     reqObj.Email,
		Password:  hashedPassword,
		Role:      database.UserRole(userRole),
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in CreateUser : %v", err))
		return
	}

	// get token
	token, err := auth.GenerateToken(theNewUser, cfg.jwt_secret)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GenerateToken : %v", err))
		return
	}

	// send response
	type respStruct struct {
		User  User   `json:"user"`
		Token string `json:"token"`
	}
	respWithJson(w, 201, respStruct{
		User: User{
			ID:        theNewUser.ID,
			CreatedAt: theNewUser.CreatedAt,
			UpdatedAt: theNewUser.UpdatedAt,
			Name:      theNewUser.Name,
			Email:     theNewUser.Email,
			Role:      string(theNewUser.Role),
		},
		Token: token,
	})
}

func (cfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {
	// get the request
	type reqstruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqstruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// check all inputs
	if reqObj.Email == "" || reqObj.Password == "" || len(reqObj.Password) < 6 {
		respWithError(w, 400, "Follow the input instructions")
		return
	}

	// check the user by email
	user, err := cfg.DB.GetUserByEmail(r.Context(), reqObj.Email)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No user with this email is registered")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in GetUserByEmail : %v", err))
			return
		}
	}

	// verify the password
	hasPasswordMatched := auth.CheckHashedPassword(user.Password, reqObj.Password)
	if !hasPasswordMatched {
		respWithError(w, 401, "password not matched")
		return
	}

	// generate token
	token, err := auth.GenerateToken(user, cfg.jwt_secret)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GenerateToken : %v", err))
		return
	}

	// send response
	type respStruct struct {
		User  User   `json:"user"`
		Token string `json:"token"`
	}
	respWithJson(w, 201, respStruct{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
		},
		Token: token,
	})

}
