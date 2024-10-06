package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bookofshame/bookofshame/internal/auth"
	"github.com/bookofshame/bookofshame/pkg/captcha"
	"github.com/bookofshame/bookofshame/pkg/jwt"
	"github.com/bookofshame/bookofshame/pkg/render"
	"github.com/go-chi/httprate"
)

type AuthHandler struct {
	ren         *render.Renderer
	j           *jwt.Jwt
	rateLimiter *httprate.RateLimiter
	recaptcha   *captcha.ReCaptcha

	authService auth.Service
}

func NewAuthHandler(ren *render.Renderer, j *jwt.Jwt, recaptcha *captcha.ReCaptcha, authService auth.Service) *AuthHandler {
	return &AuthHandler{
		ren:         ren,
		j:           j,
		rateLimiter: httprate.NewRateLimiter(5, time.Minute), // Rate-limit login at 5 req/min
		recaptcha:   recaptcha,

		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	dto := auth.UserLogin{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&dto); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if h.rateLimiter.RespondOnLimit(w, r, dto.Phone) {
		return
	}

	if err := h.recaptcha.Verify(dto.CaptchaToken); err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authService.Login(dto)
	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.j.Token(jwt.Payload{UserId: user.Id, UserLocale: user.Locale})
	if err != nil {
		h.ren.RenderJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.ren.RenderJSON(w, http.StatusOK, auth.Token{
		AccessToken: token,
	})
}

func (h *AuthHandler) Unauthorized(w http.ResponseWriter, _ *http.Request) {
	h.ren.RenderJSON(w, http.StatusUnauthorized, nil)
}
