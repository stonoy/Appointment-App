package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/Appointment-App/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

type Doctor struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Specialty     string    `json:"specialty"`
	LicenseNumber string    `json:"licence_number"`
	UserID        uuid.UUID `json:"user_id"`
}

type Availability struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Location       string    `json:"location"`
	Timing         time.Time `json:"timing"`
	Duration       int32     `json:"duration"`
	MaxPatient     int32     `json:"max_patient"`
	CurrentPatient int32     `json:"current_patient"`
	Treatment      string    `json:"treatment"`
	DoctorID       uuid.UUID `json:"doctor_id"`
}

type AppointmentResp struct {
	ID             uuid.UUID                  `json:"id"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	Status         database.AppointmentStatus `json:"status"`
	PatientID      uuid.UUID                  `json:"patient_id"`
	AvailabilityID uuid.UUID                  `json:"availability_id"`
	Location       string                     `json:"location"`
	Timing         time.Time                  `json:"timing"`
	Duration       int32                      `json:"duration"`
	Treatment      string                     `json:"treatment"`
	Name           string                     `json:"doctor_name"`
	Specialty      string                     `json:"specialty"`
}

func doctorDbToResp(doctors []database.Doctor) []Doctor {
	final := []Doctor{}

	for _, doctor := range doctors {
		final = append(final, Doctor{
			ID:            doctor.ID,
			CreatedAt:     doctor.CreatedAt,
			UpdatedAt:     doctor.UpdatedAt,
			Name:          doctor.Name,
			Specialty:     doctor.Specialty,
			LicenseNumber: doctor.LicenseNumber,
			UserID:        doctor.UserID,
		})
	}

	return final
}

func availabilityDbtoResp(availabilities []database.Availability) []Availability {
	final := []Availability{}

	for _, availability := range availabilities {
		final = append(final, Availability{
			ID:             availability.ID,
			CreatedAt:      availability.CreatedAt,
			UpdatedAt:      availability.UpdatedAt,
			Location:       availability.Location,
			Timing:         availability.Timing,
			Duration:       availability.Duration,
			MaxPatient:     availability.MaxPatient,
			CurrentPatient: availability.CurrentPatient,
			Treatment:      availability.Treatment,
			DoctorID:       availability.DoctorID,
		})
	}

	return final
}

func appointDbtoResp(appointments []database.GetAppointmentsPatientRow) []AppointmentResp {
	final := []AppointmentResp{}

	for _, appointment := range appointments {
		final = append(final, AppointmentResp{
			ID:             appointment.ID,
			CreatedAt:      appointment.CreatedAt,
			UpdatedAt:      appointment.UpdatedAt,
			Status:         appointment.Status,
			Timing:         appointment.Timing,
			Duration:       appointment.Duration,
			PatientID:      appointment.PatientID,
			AvailabilityID: appointment.AvailabilityID,
			Name:           appointment.Name,
			Location:       appointment.Location,
			Specialty:      appointment.Specialty,
			Treatment:      appointment.Treatment,
		})
	}

	return final
}
