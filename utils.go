package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("No Authorization provided")
	}

	headerSlice := strings.Fields(header)

	if len(headerSlice) < 2 || headerSlice[0] != "Bearer" {
		return "", fmt.Errorf("check auth header")
	}

	return headerSlice[1], nil
}

func GetUuidFromStr(idStr string) (uuid.UUID, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error in parsing uuid")
	}

	return id, nil
}

func GetTimeFromStr(timeStr string) (time.Time, error) {

	// Define the layout for the date-time string
	layout := time.RFC3339

	theTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return theTime.UTC(), nil
}

func ConvertInt32FromStr(str string) (int32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return int32(num), nil
}
