package providers

import (
	"github.com/stnss/dealls-interview/internal/appctx"
)

type Provider struct {
}

func NewProvider(cfg *appctx.Config) *Provider {
	return &Provider{}
}
