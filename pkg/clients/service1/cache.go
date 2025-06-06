// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package service1

import (
	"context"
	"time"
)

type cache interface {
	SetTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error)
	GetTTL(ctx context.Context, key string, value interface{}) (createdAt time.Time, ttl time.Duration, err error)
}
