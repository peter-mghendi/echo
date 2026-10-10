package export_dispatch_guest_handler

import (
	"slices"
	"testing"

	"wit_component/dispatch_guest_handler"

	witTypes "go.bytecodealliance.org/pkg/wit/types"
)

func header(k, v string) witTypes.Tuple2[string, string] {
	return witTypes.Tuple2[string, string]{F0: k, F1: v}
}

func TestHandleReturnsStatusOK(t *testing.T) {
	resp := Handle(dispatch_guest_handler.Request{})
	if resp.Status != 200 {
		t.Errorf("Status = %d, want 200", resp.Status)
	}
}

func TestHandleEchoesHeaders(t *testing.T) {
	tests := []struct {
		name    string
		headers []witTypes.Tuple2[string, string]
	}{
		{"nil", nil},
		{"empty", []witTypes.Tuple2[string, string]{}},
		{"single", []witTypes.Tuple2[string, string]{header("content-type", "text/plain")}},
		{"multiple, order preserved", []witTypes.Tuple2[string, string]{
			header("a", "1"),
			header("b", "2"),
			header("c", "3"),
		}},
		{"duplicate keys", []witTypes.Tuple2[string, string]{
			header("set-cookie", "x=1"),
			header("set-cookie", "y=2"),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Handle(dispatch_guest_handler.Request{Headers: tt.headers})
			if !slices.Equal(resp.Headers, tt.headers) {
				t.Errorf("Headers = %v, want %v", resp.Headers, tt.headers)
			}
		})
	}
}

func TestHandleEchoesBody(t *testing.T) {
	tests := []struct {
		name string
		body []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"text", []byte("hello, world")},
		{"binary", []byte{0x00, 0xff, 0x10, 0x80}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Handle(dispatch_guest_handler.Request{Body: tt.body})
			if !slices.Equal(resp.Body, tt.body) {
				t.Errorf("Body = %v, want %v", resp.Body, tt.body)
			}
		})
	}
}

func TestHandleBodyDoesNotAliasRequest(t *testing.T) {
	reqBody := []byte("original")
	resp := Handle(dispatch_guest_handler.Request{Body: reqBody})

	reqBody[0] = 'X'
	if string(resp.Body) != "original" {
		t.Errorf("mutating the request body changed the response body: got %q", resp.Body)
	}

	resp.Body[1] = 'Y'
	if string(reqBody) != "Xriginal" {
		t.Errorf("mutating the response body changed the request body: got %q", reqBody)
	}
}

func TestHandleIgnoresMethodAndPath(t *testing.T) {
	headers := []witTypes.Tuple2[string, string]{header("k", "v")}
	body := []byte("payload")

	want := Handle(dispatch_guest_handler.Request{Headers: headers, Body: body})
	for _, method := range []string{"GET", "POST", "DELETE", ""} {
		got := Handle(dispatch_guest_handler.Request{
			Method:  method,
			Path:    "/some/path",
			Headers: headers,
			Body:    body,
		})
		if got.Status != want.Status ||
			!slices.Equal(got.Headers, want.Headers) ||
			!slices.Equal(got.Body, want.Body) {
			t.Errorf("method %q: response differs from baseline", method)
		}
	}
}
