package handler

import (
	"net/http"

	"github.com/bookofshame/bookofshame/internal/auth"
	"github.com/bookofshame/bookofshame/internal/gender"
	"github.com/bookofshame/bookofshame/internal/location"
	"github.com/bookofshame/bookofshame/internal/offender"
	"github.com/bookofshame/bookofshame/internal/user"
	"github.com/bookofshame/bookofshame/pkg/captcha"
	"github.com/bookofshame/bookofshame/pkg/jwt"
	"github.com/bookofshame/bookofshame/pkg/locale"
	"github.com/bookofshame/bookofshame/pkg/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRoutes(
	j *jwt.Jwt,
	recaptcha *captcha.ReCaptcha,
	genderService gender.Service,
	locationService location.Service,
	authService auth.Service,
	userService user.Service,
	offenderService offender.Service,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.RedirectSlashes)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Compress(5))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(j.SetContext)
	mux.Use(locale.SetContext) // This must be after j.SetContext

	mux.Handle("/asset/*", http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset/"))))

	cors.AllowAll().Handler(mux)

	ren := render.NewRenderer()

	homeHandler := NewHomeHandler(ren, offenderService)
	authHandler := NewAuthHandler(ren, j, recaptcha, authService)
	locationHandler := NewLocationHandler(ren, locationService)
	userHandler := NewUserHandler(ren, j, userService)
	offenderHandler := NewOffenderHandler(ren, offenderService, genderService, locationService)

	mux.Get("/", homeHandler.Home)

	mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Get("/unauthorized", authHandler.Unauthorized)
	})

	mux.Route("/location", func(r chi.Router) {
		r.Get("/divisions", locationHandler.GetDivisions)
		r.Get("/districts", locationHandler.GetDistricts)
		r.Get("/upazilas", locationHandler.GetUpazilas)
		r.Get("/unions", locationHandler.GetUnions)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/activate", userHandler.ActivateByPhone)
		r.Post("/otp/resend", userHandler.ResendOtp)

		r.Group(func(r chi.Router) {
			r.Use(jwt.Verify)
			r.Get("/me", userHandler.Me)
		})
	})

	mux.Route("/offenders", func(r chi.Router) {
		r.Get("/", offenderHandler.GetAll)
		r.Get("/create", offenderHandler.CreateForm)
		r.Post("/", offenderHandler.Create)
		r.Delete("/{id}", offenderHandler.Delete)
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		ren.RenderJSON(w, http.StatusNotFound, nil)
	})

	return mux
}
