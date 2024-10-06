package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bookofshame/bookofshame/internal/user"
	"github.com/bookofshame/bookofshame/pkg/jwt"
	"github.com/bookofshame/bookofshame/pkg/render"
	"github.com/go-chi/httprate"
)

type UserHandler struct {
	ren         *render.Renderer
	j           *jwt.Jwt
	rateLimiter *httprate.RateLimiter

	userService user.Service
}

func NewUserHandler(ren *render.Renderer, j *jwt.Jwt, userService user.Service) *UserHandler {
	return &UserHandler{
		ren:         ren,
		rateLimiter: httprate.NewRateLimiter(1, time.Minute), // Rate-limit resend otp at 1 req/min
		j:           j,

		userService: userService,
	}
}

func (h UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	data, _ := jwt.GetDataFromContext(r.Context())
	usr, _ := h.userService.Get(data.UserId)

	h.ren.RenderJSON(w, http.StatusOK, usr)
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto := user.User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&dto); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.Create(dto); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.ren.RenderJSON(w, http.StatusOK, nil)
}

func (h UserHandler) ResendOtp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	phone := r.FormValue("phone")

	// Rate-limit login at 1 req/min.
	if h.rateLimiter.OnLimit(w, r, phone) {
		h.ren.RenderJSON(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	if err := h.userService.ResendOtp(phone); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err)
		return
	}

	h.ren.RenderJSON(w, http.StatusBadRequest, nil)
}

func (h UserHandler) ActivateByPhone(w http.ResponseWriter, r *http.Request) {
	dto := user.Otp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&dto); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.VerifyPhone(dto.Code, dto.Phone); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.ren.RenderJSON(w, http.StatusOK, nil)
}

func (h UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	activationCode := r.URL.Query().Get("code")

	if err := h.userService.VerifyEmail(activationCode); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, fmt.Errorf("%v", err))
		return
	}

	h.ren.RenderJSON(w, http.StatusOK, nil)
}
