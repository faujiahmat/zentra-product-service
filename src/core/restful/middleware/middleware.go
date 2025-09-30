package middleware

import (
	"github.com/faujiahmat/zentra-product-service/src/core/restful/client"
)

type Middleware struct {
	restfulClient *client.Restful
}

func New(rc *client.Restful) *Middleware {
	return &Middleware{
		restfulClient: rc,
	}
}
