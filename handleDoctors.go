package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/Appointment-App/internal/database"
)

func (cfg *apiConfig) createAvailability(w http.ResponseWriter, r *http.Request, doctor database.Doctor) {
	// get all the inputs
	type reqstruct struct {
		Location       string `json:"location"`
		Timing         string `json:"timing"`
		Duration       int32  `json:"duration"`
		MaxPatient     int32  `json:"max_patient"`
		CurrentPatient int32  `json:"current_patient"`
		Treatment      string `json:"treatment"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqstruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// check all inputs
	if reqObj.Location == "" || reqObj.Timing == "" || reqObj.Duration == 0 || reqObj.MaxPatient == 0 || reqObj.CurrentPatient == 0 || reqObj.Treatment == "" {
		respWithError(w, 400, "Follow the input instructions")
		return
	}

	// parse str to time.Time
	parsedTime, err := GetTimeFromStr(reqObj.Timing)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error in GetTimeFromStr : %v", err))
		return
	}

	// time is not from past
	if parsedTime.Before(time.Now()) {
		respWithError(w, 400, "can not allow past time")
		return
	}

	// create availability
	availability, err := cfg.DB.CreateAvailability(r.Context(), database.CreateAvailabilityParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Location:       reqObj.Location,
		Treatment:      reqObj.Treatment,
		MaxPatient:     reqObj.MaxPatient,
		CurrentPatient: reqObj.CurrentPatient,
		Timing:         parsedTime,
		Duration:       reqObj.Duration,
		DoctorID:       doctor.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in CreateAvailability : %v", err))
		return
	}

	// send response
	type respStruct struct {
		Msg string `json:"msg"`
	}

	respWithJson(w, 201, respStruct{
		Msg: fmt.Sprintf("Availability created with id : %v", availability.ID),
	})
}

func (cfg *apiConfig) getAvailabilityDoctor(w http.ResponseWriter, r *http.Request, doctor database.Doctor) {
	// get the query params
	queryParams := r.URL.Query()

	locationQ := queryParams.Get("location")
	start_time_Q := queryParams.Get("start_time")
	end_time_Q := queryParams.Get("end_time")
	pageQ := queryParams.Get("page")

	// initiate the default filters
	location := "%%"
	start_time := time.Now().UTC()
	end_time := time.Now().AddDate(5, 0, 0).UTC()
	var page int32 = 1

	if locationQ != "" {
		location = "%" + locationQ + "%"
	}

	if pageQ != "" {
		pageInt32, err := ConvertInt32FromStr(pageQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("Error in ConvertInt32FromStr : %v", err))
			return
		}

		page = pageInt32
	}

	if start_time_Q != "" {
		time, err := GetTimeFromStr(start_time_Q)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("Error in GetTimeFromStr : %v", err))
			return
		}

		start_time = time
	}

	if end_time_Q != "" {
		time, err := GetTimeFromStr(end_time_Q)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("Error in GetTimeFromStr : %v", err))
			return
		}

		end_time = time
	}

	// set limit and offset
	var limit int32 = 2
	offset := limit * (page - 1)

	log.Println(location, start_time, end_time, limit, offset)

	// get all availability
	allAvailabilityDoctor, err := cfg.DB.GetAvailabilityDoctor(r.Context(), database.GetAvailabilityDoctorParams{
		Location: location,
		Timing:   start_time,
		Timing_2: end_time,
		DoctorID: doctor.ID,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetAvailabilityDoctor : %v", err))
		return
	}

	// get total num of availability
	numOfAvailabilityDoctor, err := cfg.DB.GetNumAvailabilityDoctor(r.Context(), database.GetNumAvailabilityDoctorParams{
		Location: location,
		Timing:   start_time,
		Timing_2: end_time,
		DoctorID: doctor.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetNumAvailabilityDoctor : %v", err))
		return
	}

	// set num of pages
	numOfPages := int(math.Ceil(float64(numOfAvailabilityDoctor) / float64(limit)))

	// send response
	type respStruct struct {
		Availabilities []Availability `json:"avaliabilities"`
		Page           int32          `json:"page"`
		NumOfPages     int            `json:"numOfPages"`
	}

	respWithJson(w, 200, respStruct{
		Availabilities: availabilityDbtoResp(allAvailabilityDoctor),
		Page:           page,
		NumOfPages:     numOfPages,
	})
}
