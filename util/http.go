package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type simpleResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	JSON(w, simpleResp{
		Code:    400,
		Message: "Bad request!",
	})
}

func InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	JSON(w, simpleResp{
		Code:    500,
		Message: "Internal server error!",
	})
}

func JSON(w http.ResponseWriter, jdata any) {
	j, err := json.Marshal(jdata)

	if err != nil {
		log.Println("Could not parse data!")
		InternalError(w)

		return
	}

	if string(j) == "null" {
		j = []byte("[]")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
