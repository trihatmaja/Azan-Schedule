package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	az "github.com/trihatmaja/Azan-Schedule"
)

// Handler controls request flow from client to service
type Handler struct {
	azan *az.Azan
}

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

type CustomError struct {
	Message  string
	Code     int
	HTTPCode int
}

func (c CustomError) Error() string {
	return c.Message
}

// NewHandler returns a pointer of Handler instance
func NewHandler(azan *az.Azan) *Handler {
	return &Handler{
		azan: azan,
	}
}

func ResponseSuccess(data, meta interface{}) SuccessBody {
	return SuccessBody{
		Data: data,
		Meta: meta,
	}
}

func ResponseWrite(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ResponseError(errors []error) ErrorBody {
	var (
		ce CustomError
	)

	err := errors[0]
	ce, _ = err.(CustomError)

	return ErrorBody{
		Errors: []ErrorInfo{
			ErrorInfo{
				Message: ce.Message,
				Code:    ce.Code,
			},
		},
		Meta: MetaInfo{
			HTTPStatus: ce.HTTPCode,
		},
	}
}

func writeError(w http.ResponseWriter, err error) {
	res := ResponseError([]error{err})
	ResponseWrite(w, res)
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	res := ResponseSuccess(data, Meta{HTTPStatus: 200})
	ResponseWrite(w, res)
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
	}

	var req az.ApiRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		writeError(w, err)
	}

	err = h.azan.Generate(req.Latitude, req.Longitude, req.TimeZone, req.City)
	if err != nil {
		writeError(w, err)
	}

	rsp, _ := json.Marshal(req)

	writeSuccess(w, string(rsp))
}

func (h *Handler) All(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dt, err := h.azan.GetAll()
	if err != nil {
		writeError(w, err)
	}

	retval, err := json.Marshal(dt)
	if err != nil {
		writeError(w, err)
	}

	writeSuccess(w, string(retval))
}
