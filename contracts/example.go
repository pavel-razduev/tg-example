package contracts

import (
	"context"

	"tg-example/contracts/types"
)

// @tg http-prefix=v1 jsonRPC-servers log metrics trace
// @tg tagNoOmitempty=false
type ExampleService interface {
	Get(ctx context.Context, req types.BaseRequest) (resp *types.BaseResponse, err error)
}
