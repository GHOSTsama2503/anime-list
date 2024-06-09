package auth

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func NewAuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {

	return func(ctx huma.Context, next func(huma.Context)) {
		isAuthorizationRequired := false

		for _, opScheme := range ctx.Operation().Security {
			var ok bool
			if _, ok = opScheme["auth"]; ok {
				isAuthorizationRequired = true
				break
			}
		}

		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		token, err := huma.ReadCookie(ctx, "token")

		if err != nil || len(token.Value) == 0 {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}

		next(ctx)
	}
}
