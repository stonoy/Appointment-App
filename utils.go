package main

import (
	"fmt"
	"net/http"
	"strings"

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
