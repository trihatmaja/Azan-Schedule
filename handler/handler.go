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

// NewHandler returns a pointer of Handler instance
func NewHandler(azan *az.Azan) *Handler {
	return &Handler{
		azan: azan,
	}
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	var req az.ApiRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	err = h.azan.Generate(req.Latitude, req.Longitude, req.TimeZone, req.City)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	rsp, _ := json.Marshal(req)

	writeSuccess(w, string(rsp))
}

func (h *Handler) All(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dt, err := h.azan.GetAll()
	if err != nil {
		writeError(w, err, 500)
		return
	}

	retval, err := json.Marshal(dt)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	writeSuccess(w, string(retval))
}

func (h *Handler) ByCity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dt, err := h.azan.GetByCity(params.ByName("city"))
	if err != nil {
		writeError(w, err, 500)
		return
	}

	retval, err := json.Marshal(dt)
	if err != nil {
		writeError(w, err, 500)
		return
	}

	writeSuccess(w, string(retval))
}
