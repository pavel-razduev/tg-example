package contracts

import (
	"context"

	"tg-example/contracts/types"
)

// @tg http-prefix=v1
// @tg tagNoOmitempty=false
// @tg jsonRPC-servers log metrics trace
type Service2 interface {
	Set(ctx context.Context, req types.BaseRequest) (resp *types.BaseResponse, err error)
}
