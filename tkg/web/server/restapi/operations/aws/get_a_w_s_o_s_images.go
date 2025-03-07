// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAWSOSImagesHandlerFunc turns a function with the right signature into a get a w s o s images handler
type GetAWSOSImagesHandlerFunc func(GetAWSOSImagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAWSOSImagesHandlerFunc) Handle(params GetAWSOSImagesParams) middleware.Responder {
	return fn(params)
}

// GetAWSOSImagesHandler interface for that can handle valid get a w s o s images params
type GetAWSOSImagesHandler interface {
	Handle(GetAWSOSImagesParams) middleware.Responder
}

// NewGetAWSOSImages creates a new http.Handler for the get a w s o s images operation
func NewGetAWSOSImages(ctx *middleware.Context, handler GetAWSOSImagesHandler) *GetAWSOSImages {
	return &GetAWSOSImages{Context: ctx, Handler: handler}
}

/*
GetAWSOSImages swagger:route GET /api/providers/aws/osimages aws getAWSOSImages

Retrieve AWS supported os images
*/
type GetAWSOSImages struct {
	Context *middleware.Context
	Handler GetAWSOSImagesHandler
}

func (o *GetAWSOSImages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAWSOSImagesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
