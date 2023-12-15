// Package httpio provides HTTP handlers and utilities for reading and writing
// HTTP requests and responses.
//
// The handlers are designed to be used with the standard library's http
// package. They are composable and can be chained together to form a pipeline
// of handlers. The advantage of using this package is that it provides a
// consistent way of handling HTTP requests and responses. It also provides a
// way to handle errors and log them.
//
// Unlike common "handlers return errors" approach, this package uses a
// middleware approach to handle errors. This way, the error handling is
// decentralized and the handlers are in control of the error handling.
package httpio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/inancgumus/gobyexample/bite"
)

// Handler is a [http.Handler] that allows chaining handlers.
type Handler func(w http.ResponseWriter, r *http.Request) Handler

// ServeHTTP implements the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if next := h(w, r); next != nil {
		next.ServeHTTP(w, r)
	}
}

// OK returns a nil handler that stops the [Handler] chain.
func OK(w http.ResponseWriter, r *http.Request) Handler { return nil }

// Code returns a [Handler] that writes the status code to the response.
func Code(code int, h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.WriteHeader(code)
		return h
	}
}

// Text returns a [Handler] that writes the given text to the response.
func Text(s string) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, s)
		return OK
	}
}

// JSON returns a [Handler] that writes the given value to the response as JSON.
func JSON(v any) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			slog.Log(r.Context(), slog.LevelError, "internal", "url", r.URL, "message", err)
		}
		return OK
	}
}

// DecodeJSON decodes the JSON from r into v.
func DecodeJSON(r io.Reader, v any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

// Error converts an error to an HTTP status code and writes the error message
// to the response. If the error is internal, it logs it and hides the actual
// error message from the client.
func Error(err error) Handler {
	if err == nil { // no error
		return OK
	}
	var code int
	switch {
	case errors.Is(err, bite.ErrInvalidRequest):
		code = http.StatusBadRequest
	case errors.Is(err, bite.ErrExists):
		code = http.StatusConflict
	case errors.Is(err, bite.ErrNotExist):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	return func(w http.ResponseWriter, r *http.Request) Handler {
		if code == http.StatusInternalServerError {
			slog.Log(r.Context(), slog.LevelError, "internal",
				"url", r.URL, "message", err)
			err = bite.ErrInternal
		}
		return Code(code, Text(err.Error()))
	}
}
