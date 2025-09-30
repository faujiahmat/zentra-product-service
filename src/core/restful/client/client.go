package client

import "github.com/faujiahmat/zentra-product-service/src/interface/delivery"

// this main restful client
type Restful struct {
	ImageKit delivery.ImageKitRESTful
}

func NewRestful(ikc delivery.ImageKitRESTful) *Restful {
	return &Restful{
		ImageKit: ikc,
	}
}
