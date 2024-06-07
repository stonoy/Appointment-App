package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/Appointment-App/internal/database"
)

func (cfg *apiConfig) createPatient(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the inputs
	type reqStruct struct {
		Name   string `json:"name"`
		Age    int32  `json:"age"`
		Gender string `json:"gender"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// check the inputs
	if reqObj.Name == "" || reqObj.Age == 0 || reqObj.Gender == "" {
		respWithError(w, 400, "Follow the input instructions")
		return
	}

	// create patient
	patient, err := cfg.DB.CreatePatient(r.Context(), database.CreatePatientParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      reqObj.Name,
		Age:       reqObj.Age,
		Gender:    reqObj.Gender,
		UserID:    user.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in CreatePatient : %v", err))
		return
	}

	// send response
	type respStruct struct {
		Msg string `json:"msg"`
	}

	respWithJson(w, 201, respStruct{
		Msg: fmt.Sprintf("patient created - Name: %v", patient.Name),
	})
}

func (cfg *apiConfig) createAppointment(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the inputs
	type reqStruct struct {
		Id string `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error decoding request : %v", err))
		return
	}

	// convert str to uuid
	availabilityId, err := GetUuidFromStr(reqObj.Id)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error in GetUuidFromStr : %v", err))
		return
	}

	// check user is a patient
	patient, err := cfg.DB.CheckUserIsPatient(r.Context(), user.ID)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "user has not created a patient profile")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in CheckUserIsPatient : %v", err))
			return
		}
	}

	// check the availability present and obsolate if no, updated availability counts, else return
	avaliability, err := cfg.DB.UpdateCurrentPatient(r.Context(), availabilityId)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such doctor avalability found")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in UpdateCurrentPatient : %v", err))
			return
		}
	}

	// create appointment
	appointment, err := cfg.DB.CreateAppointment(r.Context(), database.CreateAppointmentParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Status:         "scheduled",
		PatientID:      patient.ID,
		AvailabilityID: avaliability.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in CreateAppointment : %v", err))
		return
	}

	// send response
	type respStruct struct {
		Msg string `json:"msg"`
	}

	respWithJson(w, 201, respStruct{
		Msg: fmt.Sprintf("Appointment created with id : %v", appointment.ID),
	})
}

func (cfg *apiConfig) getAvailability(w http.ResponseWriter, r *http.Request) {
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
	allAvailability, err := cfg.DB.GetAvailability(r.Context(), database.GetAvailabilityParams{
		Location: location,
		Timing:   start_time,
		Timing_2: end_time,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetAvailability : %v", err))
		return
	}

	// get total num of availability
	numOfAvailability, err := cfg.DB.GetNumAvailability(r.Context(), database.GetNumAvailabilityParams{
		Location: location,
		Timing:   start_time,
		Timing_2: end_time,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetNumAvailability : %v", err))
		return
	}

	// set num of pages
	numOfPages := int(math.Ceil(float64(numOfAvailability) / float64(limit)))

	// send response
	type respStruct struct {
		Availabilities []AvailabilityPatient `json:"avaliabilities"`
		Page           int32                 `json:"page"`
		NumOfPages     int                   `json:"numOfPages"`
	}

	respWithJson(w, 200, respStruct{
		Availabilities: availabilityPatientDbtoResp(allAvailability),
		Page:           page,
		NumOfPages:     numOfPages,
	})
}

func (cfg *apiConfig) getAppointments(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the query params
	queryParams := r.URL.Query()

	statusQ := queryParams.Get("status")
	locationQ := queryParams.Get("location")
	treatmentQ := queryParams.Get("treatment")
	start_time_Q := queryParams.Get("start_time")
	end_time_Q := queryParams.Get("end_time")
	pageQ := queryParams.Get("page")

	// initiate default params
	location := "%%"
	status := "scheduled"
	treatment := "%%"
	start_time := time.Now().UTC()
	end_time := time.Now().AddDate(5, 0, 0).UTC()
	var page int32 = 1

	// set the queryparams
	if locationQ != "" {
		location = "%" + locationQ + "%"
	}

	if statusQ != "" {
		status = statusQ
	}

	if treatmentQ != "" {
		treatment = "%" + treatmentQ + "%"
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

	log.Println(status, treatment, location, start_time, end_time, limit, offset)

	// check user is a patient
	patient, err := cfg.DB.CheckUserIsPatient(r.Context(), user.ID)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "user has not created a patient profile")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in CheckUserIsPatient : %v", err))
			return
		}
	}

	// get appointments
	allAppointmentsPatient, err := cfg.DB.GetAppointmentsPatient(r.Context(), database.GetAppointmentsPatientParams{
		Status:    database.AppointmentStatus(status),
		Location:  location,
		Treatment: treatment,
		Timing:    start_time,
		Timing_2:  end_time,
		PatientID: patient.ID,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetAppointmentsPatient : %v", err))
		return
	}

	// get num of appointments
	numOfAllAppointmentPatient, err := cfg.DB.GetNumAppointmentsPatient(r.Context(), database.GetNumAppointmentsPatientParams{
		Status:    database.AppointmentStatus(status),
		Location:  location,
		Treatment: treatment,
		Timing:    start_time,
		Timing_2:  end_time,
		PatientID: patient.ID,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("Error in GetNumAppointmentsPatient : %v", err))
		return
	}

	// num of pages
	numOfPages := int(math.Ceil(float64(numOfAllAppointmentPatient) / float64(limit)))

	// send response
	type respStruct struct {
		Appointments []AppointmentResp `json:"appointments"`
		Page         int32             `json:"page"`
		NumOfPages   int               `json:"numOfPages"`
	}

	respWithJson(w, 200, respStruct{
		Appointments: appointDbtoResp(allAppointmentsPatient),
		Page:         page,
		NumOfPages:   numOfPages,
	})

}

func (cfg *apiConfig) DeleteAppointment(w http.ResponseWriter, r *http.Request, user database.User) {
	// get url param
	appointIdStr := chi.URLParam(r, "appointmentId")

	// convert id str- -> uuid
	appointId, err := GetUuidFromStr(appointIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("Error in GetUuidFromStr : %v", err))
		return
	}

	// user -> patient
	patient, err := cfg.DB.CheckUserIsPatient(r.Context(), user.ID)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "user has not created a patient profile")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in CheckUserIsPatient : %v", err))
			return
		}
	}

	// delete appointment
	appointment, err := cfg.DB.DeleteAppointment(r.Context(), database.DeleteAppointmentParams{
		PatientID: patient.ID,
		ID:        appointId,
	})
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such appointment found")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in DeleteAppointment : %v", err))
			return
		}
	}

	// update availability current_user count
	avail, err := cfg.DB.DelCurrentPatient(r.Context(), appointment.AvailabilityID)
	if err != nil {
		// no such records found
		if err == sql.ErrNoRows {
			respWithError(w, 400, "No such valid availablity found with this appointment")
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("Error in DeleteAppointment : %v", err))
			return
		}
	}

	log.Println(avail.CurrentPatient)

	// send response
	type respStruct struct {
		Msg string `json:"msg"`
	}

	respWithJson(w, 200, respStruct{
		Msg: fmt.Sprintf("Appointment with id : %v, Deleted", appointment.ID),
	})
}
