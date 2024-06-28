package middlewares

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/danielgtaylor/huma/v2"
	"github.com/ghostsama2503/anime-list/services/sessions"
)

func AuthForbbiden(api huma.API, ctx huma.Context) {
	huma.WriteErr(api, ctx, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

func Auth(api huma.API, sessionsService *sessions.SessionsService) HumaMiddleware {
	return func(ctx huma.Context, next func(huma.Context)) {

		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			var ok bool
			if _, ok = opScheme["cookie"]; ok {
				isAuthorizationRequired = true
				break
			}
		}

		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		session, err := huma.ReadCookie(ctx, "session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				AuthForbbiden(api, ctx)
				return
			}

			log.Error("error reading session cookie", "err", err)
			return
		}

		sessionCookie, err := sessionsService.ParseSessionCookie(session.Value)
		if err != nil {
			AuthForbbiden(api, ctx)
			return
		}

		isValidSession, err := sessionsService.Validate(ctx.Context(), sessionCookie.SessionId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				AuthForbbiden(api, ctx)
				return
			}

			errMsg := "error validating session"
			log.Error(errMsg, "err", err)

			huma.WriteErr(api, ctx, http.StatusInternalServerError, errMsg)
			return
		}

		if !isValidSession {
			AuthForbbiden(api, ctx)
			return
		}

		next(ctx)
	}
}
