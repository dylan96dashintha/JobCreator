package authentication

import (
	"context"
	"errors"
	"github.com/JobCreator/app/http/globals"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v5"
	logger "github.com/sirupsen/logrus"
)

func AuthenticateHandler(key []byte) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			tokenString, ok := ctx.Value(globals.JWTTokenContextKey).(string)
			if !ok {
				logger.Error(ctx.Value(globals.UUID), "token not found")
				return response, errors.New("no token found")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				_, ok = token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					logger.Error(ctx.Value(globals.UUID), "error in decoding token")
					return response, errors.New("unexpected signing method")
				}
				return key, nil
			})

			if err != nil {
				logger.Error(ctx.Value(globals.UUID), "Invalid token ", err)
				return response, errors.New("invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				logger.Error(ctx.Value(globals.UUID), "invalid token")
			} else {
				passengerId := claims["username"].(string)
				logger.Info(ctx.Value(globals.UUID), "access granted", passengerId)
			}

			return next(ctx, request)
		}
	}

}
