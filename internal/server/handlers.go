package server

import (
	"net/http"

	"weber-br/internal/drop"
)

type weberRequest struct {
	RhoGas    float64 `json:"rho_g"`
	RhoLiquid float64 `json:"rho_l"`
	MuLiquid  float64 `json:"mu_l"`
	Sigma     float64 `json:"sigma"`
	U         float64 `json:"U"`
	D         float64 `json:"d"`
}

type uCritRequest struct {
	RhoGas    float64 `json:"rho_g"`
	RhoLiquid float64 `json:"rho_l"`
	MuLiquid  float64 `json:"mu_l"`
	Sigma     float64 `json:"sigma"`
	D         float64 `json:"d"`
}

func weberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w, "POST")
		return
	}
	var req weberRequest
	if !readJSON(w, r, &req) {
		return
	}
	in := drop.Input{
		Fluids: drop.Fluids{RhoGas: req.RhoGas, RhoLiquid: req.RhoLiquid, MuLiquid: req.MuLiquid, Sigma: req.Sigma},
		U:      req.U, D: req.D,
	}
	result, err := drop.Evaluate(in)
	if err != nil {
		writeValidationOutcome(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func uCritHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w, "POST")
		return
	}
	var req uCritRequest
	if !readJSON(w, r, &req) {
		return
	}
	f := drop.Fluids{RhoGas: req.RhoGas, RhoLiquid: req.RhoLiquid, MuLiquid: req.MuLiquid, Sigma: req.Sigma}
	oh, err := drop.OhnesorgeNumber(f, req.D)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	criticalU, err := drop.CriticalVelocity(f, req.D, drop.CriticalWeber(oh))
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"U_crit": criticalU, "We_crit": drop.CriticalWeber(oh), "Oh": oh,
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w, "GET")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"name": "weber-br", "version": "1.0.0"})
}
