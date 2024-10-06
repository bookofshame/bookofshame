package handler

import (
	"net/http"
	"strconv"

	"github.com/bookofshame/bookofshame/internal/location"
	"github.com/bookofshame/bookofshame/pkg/render"
)

type LocationHandler struct {
	ren             *render.Renderer
	locationService location.Service
}

func NewLocationHandler(ren *render.Renderer, locationService location.Service) *LocationHandler {
	return &LocationHandler{
		ren:             ren,
		locationService: locationService,
	}
}

func (h LocationHandler) GetDivisions(w http.ResponseWriter, _ *http.Request) {
	divisions, err := h.locationService.GetDivisions()
	if err != nil {
		h.ren.RenderJSON(w, http.StatusInternalServerError, err.Error())
	}

	h.ren.RenderJSON(w, http.StatusOK, divisions)
}

func (h LocationHandler) GetDistricts(w http.ResponseWriter, r *http.Request) {
	divisionId := 0

	if id := r.URL.Query().Get("divisionId"); id != "" {
		divisionId, _ = strconv.Atoi(id)
	}

	districts, err := h.locationService.GetDistricts(divisionId)
	if err != nil {
		h.ren.RenderJSON(w, http.StatusInternalServerError, err.Error())
	}

	h.ren.RenderJSON(w, http.StatusOK, districts)
}

func (h LocationHandler) GetUpazilas(w http.ResponseWriter, r *http.Request) {
	districtId := 0

	if id := r.URL.Query().Get("districtId"); id != "" {
		districtId, _ = strconv.Atoi(id)
	}

	upazilas, err := h.locationService.GetUpazilas(districtId)
	if err != nil {
		h.ren.RenderJSON(w, http.StatusInternalServerError, err.Error())
	}

	h.ren.RenderJSON(w, http.StatusOK, upazilas)
}

func (h LocationHandler) GetUnions(w http.ResponseWriter, r *http.Request) {
	upazilaId := 0

	if id := r.URL.Query().Get("upazilaId"); id != "" {
		upazilaId, _ = strconv.Atoi(id)
	}

	unions, err := h.locationService.GetUnions(upazilaId)
	if err != nil {
		h.ren.RenderJSON(w, http.StatusInternalServerError, err.Error())
	}

	h.ren.RenderJSON(w, http.StatusOK, unions)
}
