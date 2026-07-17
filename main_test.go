package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfileHandlerIsIsolated(t *testing.T) {
	handler := profileHandler()

	profile := httptest.NewRecorder()
	handler.ServeHTTP(profile, httptest.NewRequest(http.MethodGet, "/debug/pprof/", nil))
	if profile.Code != http.StatusOK {
		t.Fatalf("pprof index status = %d, want %d", profile.Code, http.StatusOK)
	}

	root := httptest.NewRecorder()
	handler.ServeHTTP(root, httptest.NewRequest(http.MethodGet, "/", nil))
	if root.Code != http.StatusNotFound {
		t.Fatalf("unrelated root status = %d, want %d", root.Code, http.StatusNotFound)
	}
}
