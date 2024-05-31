package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respWithError(w http.ResponseWriter, code int, msg string) {
	type respStruct struct {
		Msg string `json:"msg"`
	}

	if code >= 500 {
		log.Printf("Internal Server Error : %v", msg)
	}

	respWithJson(w, code, respStruct{Msg: msg})
}

func respWithJson(w http.ResponseWriter, code int, response interface{}) {
	// convert the response which can of any type to byte
	dat, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Can not marshal response to byte"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
