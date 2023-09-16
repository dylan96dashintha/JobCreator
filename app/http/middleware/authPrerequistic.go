package middleware

import (
	"context"
	"github.com/JobCreator/app/http/globals"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"strings"
)

func HttpToContext() kithttp.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {

		ct := context.Background()
		registrationSource := req.Header.Get("source")
		if len(strings.TrimSpace(registrationSource)) == 0 {
			registrationSource = ""
		}

		ct = context.WithValue(ct, globals.SOURCE_KEY, registrationSource)

		token, ok := extractTokenFromAuthHeader(req.Header.Get("Authorization"))
		if !ok {
			log.WithPrefix(log.NewNopLogger(), "No authentication")
		} else {
			ct = context.WithValue(ct, globals.JWTTokenContextKey, token)
		}
		return ct

	}
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != globals.Bearer {
		return "", false
	}

	return authHeaderParts[1], true
}
