package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/Appointment-App/internal/database"
)

func (cfg *apiConfig) createDoctors(w http.ResponseWriter, r *http.Request) {
	// get the request
	type reqStruct struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Specialty      string `json:"specialty"`
		License_Number string `json:"licence_number"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// check all inputs
	if reqObj.Email == "" || reqObj.Name == "" || reqObj.License_Number == "" || reqObj.Specialty == "" {
		respWithError(w, 400, "Follow the input instructions")
		return
	}

	// get the user and update role to doctor from email
	user, err := cfg.DB.UpdateToDoctor(r.Context(), database.UpdateToDoctorParams{
		Role:  "doctor",
		Email: reqObj.Email,
	})
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No user with this email is registered")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in UpdateToDoctor : %v", err))
			return
		}
	}

	// create doctor
	doctor, err := cfg.DB.CreateDoctor(r.Context(), database.CreateDoctorParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		Name:          reqObj.Name,
		Specialty:     reqObj.Specialty,
		LicenseNumber: reqObj.License_Number,
		UserID:        user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in CreateDoctor : %v", err))
		return
	}

	// send response
	type respStruct struct {
		Msg string `json:"msg"`
	}

	respWithJson(w, 201, respStruct{
		Msg: fmt.Sprintf("doctor created - Name: %v, Licence Number: %v", doctor.Name, doctor.LicenseNumber),
	})
}

func (cfg *apiConfig) getAllDoctors(w http.ResponseWriter, r *http.Request) {
	// get the query params
	queryParams := r.URL.Query()

	nameQ := queryParams.Get("name")
	specialtyQ := queryParams.Get("specialty")
	pageQ := queryParams.Get("page")

	// initiate a search value
	name := "%%"
	specialty := "%%"
	var page int32 = 1

	// modify the search value according to query params
	if nameQ != "" {
		name = "%" + nameQ + "%"
	}

	if specialtyQ != "" {
		specialty = "%" + specialtyQ + "%"
	}

	if pageQ != "" {
		pageInt32, err := ConvertInt32FromStr(pageQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("Error in ConvertInt32FromStr : %v", err))
			return
		}

		page = pageInt32
	}

	// set limit and offset
	var limit int32 = 2
	offset := limit * (page - 1)

	// get doctors
	doctors, err := cfg.DB.GetAllDoctorsWithPage(r.Context(), database.GetAllDoctorsWithPageParams{
		Specialty: specialty,
		Name:      name,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetAllDoctorsWithPage : %v", err))
		return
	}

	// get number of doctors
	numOfDoctors, err := cfg.DB.GetNumOfAllDoctorsWithPage(r.Context(), database.GetNumOfAllDoctorsWithPageParams{
		Specialty: specialty,
		Name:      name,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetNumOfAllDoctorsWithPage : %v", err))
		return
	}

	// calculate total number of pages
	numOfPages := int(math.Ceil(float64(numOfDoctors) / float64(limit)))

	// send response
	type respStruct struct {
		Doctors    []Doctor `json:"doctors"`
		Page       int32    `json:"page"`
		NumOfPages int      `json:"numOfPages"`
	}

	respWithJson(w, 200, respStruct{
		Doctors:    doctorDbToResp(doctors),
		Page:       page,
		NumOfPages: numOfPages,
	})
}
