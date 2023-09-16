package domain

import "context"

type WrappedResponse struct {
	Response interface{}
	Ctx      context.Context
}
