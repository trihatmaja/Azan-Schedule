package handler

import (
	"encoding/json"
	"net/http"
)

type SuccessBody struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

// Meta is used to consolidate all meta statuses
type Meta struct {
	HTTPStatus int `json:"http_status"`
}

type MetaInfo struct {
	HTTPStatus int `json:"http_status"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ResponseSuccess(data, meta interface{}) SuccessBody {
	return SuccessBody{
		Data: data,
		Meta: meta,
	}
}

func ResponseWrite(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}

func ResponseError(err error, httpstatus int) ErrorBody {

	return ErrorBody{
		Errors: []ErrorInfo{
			ErrorInfo{
				Message: err.Error(),
				Code:    0,
			},
		},
		Meta: MetaInfo{
			HTTPStatus: httpstatus,
		},
	}
}

func writeError(w http.ResponseWriter, err error, httpstatus int) {
	res := ResponseError(err, httpstatus)
	ResponseWrite(w, res)
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	res := ResponseSuccess(data, Meta{HTTPStatus: 200})
	ResponseWrite(w, res)
}
