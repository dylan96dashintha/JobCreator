package endpoints

import (
	"context"
	"github.com/JobCreator/app/http/domain"
	"github.com/go-kit/kit/endpoint"
)

func CreateJob() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		//response, err = svc.GetDriverSupportReasons(ctx, input)
		return domain.WrappedResponse{
			Response: response,
			Ctx:      ctx,
		}, err
	}
}
