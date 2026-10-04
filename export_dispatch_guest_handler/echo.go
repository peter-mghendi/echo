package export_dispatch_guest_handler

import (
	"slices"
	"wit_component/dispatch_guest_types"
)

func Handle(req dispatch_guest_types.Request) dispatch_guest_types.Response {
	return dispatch_guest_types.Response{
		Status:  200,
		Headers: req.Headers,
		Body:    slices.Clone(req.Body),
	}
}
