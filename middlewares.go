package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/stonoy/Appointment-App/auth"
	"github.com/stonoy/Appointment-App/internal/database"
)

type requiredFuncType func(w http.ResponseWriter, r *http.Request, user database.User)

// any authenticated user
func (cfg *apiConfig) onlyForAuthinticatedUser(givenFunc requiredFuncType) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from authorization header
		token, err := GetTokenFromHeader(r)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("Error in GetTokenFromHeader : %v", err))
			return
		}

		// verify and extract token
		_, idStr, err := auth.CheckAndExtractToken(token, cfg.jwt_secret)
		if err != nil {
			respWithError(w, 403, fmt.Sprintf("Error in CheckAndExtractToken : %v", err))
			return
		}

		// parse the id
		id, err := GetUuidFromStr(idStr)
		if err != nil {
			respWithError(w, 500, fmt.Sprintf("Error in GetUuidFromStr : %v", err))
			return
		}

		// get user from userid
		user, err := cfg.DB.GetUserById(r.Context(), id)
		if err != nil {
			// no such records found
			if err == sql.ErrNoRows {
				respWithError(w, 400, "No such user exist")
				return
			} else {
				respWithError(w, 500, fmt.Sprintf("Error in GetUserById : %v", err))
				return
			}
		}

		// call givenFunc
		givenFunc(w, r, user)
	}
}

type requiredFuncTypeForDoctor func(w http.ResponseWriter, r *http.Request, doctor database.Doctor)

// any doctor
func (cfg *apiConfig) onlyForDoctor(givenFunc requiredFuncTypeForDoctor) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from authorization header
		token, err := GetTokenFromHeader(r)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("Error in GetTokenFromHeader : %v", err))
			return
		}

		// verify and extract token
		role, idStr, err := auth.CheckAndExtractToken(token, cfg.jwt_secret)
		if err != nil {
			respWithError(w, 403, fmt.Sprintf("Error in CheckAndExtractToken : %v", err))
			return
		}

		// check role and only allow doctor and admin
		if role != "doctor" {
			respWithError(w, 403, "Not Authorised")
			return
		}

		// parse the id
		id, err := GetUuidFromStr(idStr)
		if err != nil {
			respWithError(w, 500, fmt.Sprintf("Error in GetUuidFromStr : %v", err))
			return
		}

		// get doctor from userid
		doctor, err := cfg.DB.GetDoctorFromUserId(r.Context(), id)
		if err != nil {
			// no such records found
			if err == sql.ErrNoRows {
				respWithError(w, 400, "No such doctor exist")
				return
			} else {
				respWithError(w, 500, fmt.Sprintf("Error in GetDoctorFromUserId : %v", err))
				return
			}
		}

		// call givenFunc
		givenFunc(w, r, doctor)
	}
}

// one admin
func (cfg *apiConfig) onlyForAdmin(givenFunc func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from authorization header
		token, err := GetTokenFromHeader(r)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("Error in GetTokenFromHeader : %v", err))
			return
		}

		// verify and extract token
		role, _, err := auth.CheckAndExtractToken(token, cfg.jwt_secret)
		if err != nil {
			respWithError(w, 403, fmt.Sprintf("Error in CheckAndExtractToken : %v", err))
			return
		}

		// check role and only allow doctor and admin
		if role != "admin" {
			respWithError(w, 403, "Not Authorised")
			return
		}

		// // parse the id
		// id, err := GetUuidFromStr(idStr)
		// if err != nil {
		// 	respWithError(w, 500, fmt.Sprintf("Error in GetUuidFromStr : %v", err))
		// 	return
		// }

		// // get user from userid
		// user, err := cfg.DB.GetUserById(r.Context(), id)
		// if err != nil {
		// 	// no such records found
		// 	if err == sql.ErrNoRows {
		// 		respWithError(w, 400, "No such user exist")
		// 		return
		// 	} else {
		// 		respWithError(w, 500, fmt.Sprintf("Error in GetUserById : %v", err))
		// 		return
		// 	}
		// }

		// call givenFunc
		givenFunc(w, r)
	}
}
