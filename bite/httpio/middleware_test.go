package httpio

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	slog.SetDefault(log)

	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
	})
	handler = LoggingMiddleware(handler)
	handler.ServeHTTP(
		httptest.NewRecorder(),
		httptest.NewRequest("GET", "/test", nil),
	)

	got := buf.String()
	if !strings.Contains(got, "GET") {
		t.Error("want GET in the log")
	}
	if !strings.Contains(got, "418") {
		t.Error("want 418 in the log")
	}
	if !strings.Contains(got, "/test") {
		t.Errorf("want /test in the log")
	}
	if t.Failed() {
		t.Log("got:", got)
	}
}

func TestTraceMiddleware(t *testing.T) {
	var (
		traceID int64
		ok      bool
	)

	var handler http.Handler
	handler = http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		traceID, ok = TraceID(r.Context())
	})
	handler = TraceMiddleware(handler)

	handler.ServeHTTP(
		httptest.NewRecorder(),
		httptest.NewRequest("GET", "/test", nil),
	)
	if !ok {
		t.Fatal("got context without a trace id")
	}
	if traceID <= 0 {
		t.Fatalf("got %d, want a positive trace id", traceID)
	}

	prev := traceID
	handler.ServeHTTP(
		httptest.NewRecorder(),
		httptest.NewRequest("GET", "/test", nil),
	)
	if prev == traceID {
		t.Fatalf("got duplicate trace id: %d", traceID)
	}
}

func TestLogHandler(t *testing.T) {
	var buf bytes.Buffer

	ctx := SetTraceID(context.Background(), 42)
	log := slog.New(&LogHandler{
		Handler: slog.NewTextHandler(&buf, nil),
	})

	log.Log(ctx, slog.LevelInfo, "test")

	if got := buf.String(); !strings.Contains(got, "42") {
		t.Errorf("want trace id %d in the log:%s", 42, got)
	}
}
