// Code generated by goa v3.11.0, DO NOT EDIT.
//
// locator endpoints
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/locator/design -o
// services/locator

package locator

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "locator" service endpoints.
type Endpoints struct {
	GetLocation goa.Endpoint
}

// NewEndpoints wraps the methods of the "locator" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetLocation: NewGetLocationEndpoint(s),
	}
}

// Use applies the given middleware to all the "locator" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetLocation = m(e.GetLocation)
}

// NewGetLocationEndpoint returns an endpoint function that calls the method
// "get_location" of service "locator".
func NewGetLocationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(string)
		return s.GetLocation(ctx, p)
	}
}
