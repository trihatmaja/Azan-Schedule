package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	writeSuccess(w, req)
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "generate", "ok", elapsedTime)
}

func (h *Handler) ByCity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		an := az.CalcResult{}
		json.Unmarshal(c, &an)
		writeSuccess(w, an)
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

	writeSuccess(w, dt)
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}

func (h *Handler) ByCityDate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		an := az.CalcResult{}
		json.Unmarshal(c, &an)
		writeSuccess(w, an)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "ok", elapsedTime)
		TraceRequestTimeCache(r.Method, "all", "HIT", elapsedTime)
		return
	}

	tm, err := time.Parse("20060102", params.ByName("date"))
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}
	dt, err := h.azan.GetByCityDate(params.ByName("city"), tm)
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

	writeSuccess(w, dt)
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}

func (h *Handler) ByCityMonth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		an := az.CalcResult{}
		json.Unmarshal(c, &an)
		writeSuccess(w, an)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "ok", elapsedTime)
		TraceRequestTimeCache(r.Method, "all", "HIT", elapsedTime)
		return
	}

	i, err := strconv.Atoi(params.ByName("month"))
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	dt, err := h.azan.GetByCityMonth(params.ByName("city"), i)
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

	writeSuccess(w, dt)
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}

func (h *Handler) ByCities(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	startTime := time.Now()

	key := r.RequestURI

	c, err := h.azan.GetCache(key)
	if err == nil {
		an := []az.CalcResult{}
		json.Unmarshal(c, &an)
		writeSuccess(w, an)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "ok", elapsedTime)
		TraceRequestTimeCache(r.Method, "all", "HIT", elapsedTime)
		return
	}

	dt, err := h.azan.GetCities()
	if err != nil {
		writeError(w, err, 500)
		elapsedTime := time.Since(startTime).Seconds()
		TraceRequestTime(r.Method, "all", "fail", elapsedTime)
		return
	}

	go func() {
		retval, err := json.Marshal(dt)
		if err != nil {
			writeError(w, err, 500)
			elapsedTime := time.Since(startTime).Seconds()
			TraceRequestTime(r.Method, "all", "fail", elapsedTime)
			return
		}

		h.azan.SetCache(key, retval)
	}()

	writeSuccess(w, dt)
	elapsedTime := time.Since(startTime).Seconds()
	TraceRequestTime(r.Method, "all", "ok", elapsedTime)
	TraceRequestTimeCache(r.Method, "all", "MISS", elapsedTime)
}
