package export_dispatch_guest_handler

import (
	"slices"
	"wit_component/dispatch_guest_handler"
)

func Handle(req dispatch_guest_handler.Request) dispatch_guest_handler.Response {
	return dispatch_guest_handler.Response{
		Status:  200,
		Headers: req.Headers,
		Body:    slices.Clone(req.Body),
	}
}
