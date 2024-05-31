package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/stonoy/Appointment-App/auth"
	"github.com/stonoy/Appointment-App/internal/database"
)

type requiredFuncType func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) onlyForAuthinticatedUser(givenFunc requiredFuncType) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from authorization header
		token, err := GetTokenFromHeader(r)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("Error in GetTokenFromHeader : %v", err))
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
