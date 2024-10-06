package handler

import (
	"net/http"
	"strconv"

	"github.com/bookofshame/bookofshame/internal/gender"
	"github.com/bookofshame/bookofshame/internal/location"
	"github.com/bookofshame/bookofshame/internal/offender"
	"github.com/bookofshame/bookofshame/pkg/render"
	"github.com/go-chi/chi/v5"
)

type OffenderHandler struct {
	ren             *render.Renderer
	offenderService offender.Service
	genderService   gender.Service
	locationService location.Service
}

func NewOffenderHandler(ren *render.Renderer, offenderService offender.Service, genderService gender.Service, locationService location.Service) *OffenderHandler {
	return &OffenderHandler{
		ren:             ren,
		offenderService: offenderService,
		genderService:   genderService,
		locationService: locationService,
	}
}

func (h OffenderHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	offenders, _ := h.offenderService.GetAll()

	h.ren.RenderJSON(w, http.StatusOK, offenders)
}

func (h OffenderHandler) CreateForm(w http.ResponseWriter, _ *http.Request) {
	genders, _ := h.genderService.GetAll()

	h.ren.RenderJSON(w, http.StatusOK, genders)
}

func (h OffenderHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	file, _, _ := r.FormFile("photo")

	divisionId, _ := strconv.Atoi(r.FormValue("division"))
	districtId, _ := strconv.Atoi(r.FormValue("district"))
	upazilaId, _ := strconv.Atoi(r.FormValue("upazila"))
	unionId, _ := strconv.Atoi(r.FormValue("union"))

	err = h.offenderService.Create(offender.Offender{
		FullName:   r.FormValue("fullName"),
		Address:    r.FormValue("address"),
		DivisionId: divisionId,
		DistrictId: districtId,
		UpazilaId:  &upazilaId,
		UnionId:    &unionId,
		Metadata:   "{}",
	}, file)

	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
	} else {
		h.ren.RenderJSON(w, http.StatusOK, nil)
	}
}

func (h OffenderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.offenderService.Delete(id)
	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.ren.RenderJSON(w, http.StatusOK, nil)
}
