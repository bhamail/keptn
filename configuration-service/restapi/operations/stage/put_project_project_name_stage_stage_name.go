// Code generated by go-swagger; DO NOT EDIT.

package stage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutProjectProjectNameStageStageNameHandlerFunc turns a function with the right signature into a put project project name stage stage name handler
type PutProjectProjectNameStageStageNameHandlerFunc func(PutProjectProjectNameStageStageNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutProjectProjectNameStageStageNameHandlerFunc) Handle(params PutProjectProjectNameStageStageNameParams) middleware.Responder {
	return fn(params)
}

// PutProjectProjectNameStageStageNameHandler interface for that can handle valid put project project name stage stage name params
type PutProjectProjectNameStageStageNameHandler interface {
	Handle(PutProjectProjectNameStageStageNameParams) middleware.Responder
}

// NewPutProjectProjectNameStageStageName creates a new http.Handler for the put project project name stage stage name operation
func NewPutProjectProjectNameStageStageName(ctx *middleware.Context, handler PutProjectProjectNameStageStageNameHandler) *PutProjectProjectNameStageStageName {
	return &PutProjectProjectNameStageStageName{Context: ctx, Handler: handler}
}

/* PutProjectProjectNameStageStageName swagger:route PUT /project/{projectName}/stage/{stageName} Stage putProjectProjectNameStageStageName

INTERNAL Endpoint: Update the specified stage

*/
type PutProjectProjectNameStageStageName struct {
	Context *middleware.Context
	Handler PutProjectProjectNameStageStageNameHandler
}

func (o *PutProjectProjectNameStageStageName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutProjectProjectNameStageStageNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
