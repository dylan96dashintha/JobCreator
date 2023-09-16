package response

import (
	"context"
	"net/http"
)

func EncodeCreateJobResponseJson(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	return
}
