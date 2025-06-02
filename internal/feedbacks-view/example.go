package feedbacks_view

import (
	"context"

	"tg-example/internal/feedbacks-view/types"
)

// @tg http-prefix=v1 jsonRPC-servers log metrics trace
// @tg packageJSON=github.com/goccy/go-json
// @tg tagNoOmitempty=false
type ExampleService interface {
	Get(ctx context.Context, req types.BaseRequest) (resp *types.BaseResponse, err error)
}
