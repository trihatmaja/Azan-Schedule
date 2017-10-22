package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

func (h *Handler) Metrics(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	PromHandler(w, r)
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "generate", "fail", elapsedTime)
		return
	}

	var req ApiRequest

	err = json.Unmarshal(body, &req)
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "generate", "fail", elapsedTime)
		return
	}

	err = h.azan.Generate(req.Latitude, req.Longitude, req.TimeZone, req.City)
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "generate", "fail", elapsedTime)
		return
	}

	rsp, _ := json.Marshal(req)

	writeSuccess(w, string(rsp))
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "generate", "ok", elapsedTime)
}

func (h *Handler) All(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		writeSuccess(w, string(c))
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "ok", elapsedTime)
		TraceRequestTimeCache(r.Method, "all", "HIT", elapsedTime)
		return
	}

	dt, err := h.azan.GetAll()
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	retval, err := json.Marshal(dt)
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	go func() {
		h.azan.SetCache(key, retval)
	}()

	writeSuccess(w, string(retval))
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}

func (h *Handler) ByCity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		writeSuccess(w, string(c))
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "ok", elapsedTime)
		TraceRequestTimeCache(r.Method, "all", "HIT", elapsedTime)
		return
	}

	dt, err := h.azan.GetByCity(params.ByName("city"))
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	retval, err := json.Marshal(dt)
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	go func() {
		h.azan.SetCache(key, retval)
	}()

	writeSuccess(w, string(retval))
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}
