// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DownloadClusterFilesHandlerFunc turns a function with the right signature into a download cluster files handler
type DownloadClusterFilesHandlerFunc func(DownloadClusterFilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadClusterFilesHandlerFunc) Handle(params DownloadClusterFilesParams) middleware.Responder {
	return fn(params)
}

// DownloadClusterFilesHandler interface for that can handle valid download cluster files params
type DownloadClusterFilesHandler interface {
	Handle(DownloadClusterFilesParams) middleware.Responder
}

// NewDownloadClusterFiles creates a new http.Handler for the download cluster files operation
func NewDownloadClusterFiles(ctx *middleware.Context, handler DownloadClusterFilesHandler) *DownloadClusterFiles {
	return &DownloadClusterFiles{Context: ctx, Handler: handler}
}

/*DownloadClusterFiles swagger:route GET /clusters/{clusterId}/downloads/files inventory downloadClusterFiles

Download files relating to the installed/installing cluster

*/
type DownloadClusterFiles struct {
	Context *middleware.Context
	Handler DownloadClusterFilesHandler
}

func (o *DownloadClusterFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownloadClusterFilesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
