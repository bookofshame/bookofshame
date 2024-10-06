package locale

import (
	"net/http"

	"github.com/bookofshame/bookofshame/pkg/jwt"
	"github.com/invopop/ctxi18n"
)

// SetContext Middleware to set the locale in the context.
// This must be used after session middleware because it needs the session to get user's locale
func SetContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := jwt.GetDataFromContext(r.Context())
		ctx, _ := ctxi18n.WithLocale(r.Context(), data.UserLocale)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
