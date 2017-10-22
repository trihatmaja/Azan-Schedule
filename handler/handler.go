package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	az "github.com/trihatmaja/Azan-Schedule"
)

type Handler struct {
	azan *az.Azan
}

type ApiRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
	TimeZone  float64 `json:"tz"`
	City      string  `json:"city"`
}

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

	var req ApiRequest

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
	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		writeSuccess(w, string(c))
		return
	}

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

	go func() {
		h.azan.SetCache(key, retval)
	}()

	writeSuccess(w, string(retval))
}

func (h *Handler) ByCity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		writeSuccess(w, string(c))
		return
	}

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

	go func() {
		h.azan.SetCache(key, retval)
	}()

	writeSuccess(w, string(retval))
}
