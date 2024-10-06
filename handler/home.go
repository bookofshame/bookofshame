package handler

import (
	"net/http"

	"github.com/bookofshame/bookofshame/internal/offender"
	"github.com/bookofshame/bookofshame/pkg/render"
)

type HomeHandler struct {
	ren             *render.Renderer
	offenderService offender.Service
}

func NewHomeHandler(ren *render.Renderer, offenderService offender.Service) *HomeHandler {
	return &HomeHandler{
		ren:             ren,
		offenderService: offenderService,
	}
}

func (h HomeHandler) Home(w http.ResponseWriter, _ *http.Request) {
	h.ren.RenderJSON(w, http.StatusOK, struct {
	}{})
}
