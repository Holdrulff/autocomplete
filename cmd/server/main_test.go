package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildHandlerLoadsSnapshotAndServesGraphQL(t *testing.T) {
	snapshotPath := filepath.Join(
		t.TempDir(),
		"stack-overflow-tags.json",
	)

	snapshot := `{
		"source": "https://api.stackexchange.com",
		"attribution": "Stack Overflow",
		"generated_at": "2026-08-14T12:00:00Z",
		"tags": [
			{
				"name": "javascript",
				"count": 100
			}
		]
	}`

	if err := os.WriteFile(
		snapshotPath,
		[]byte(snapshot),
		0o600,
	); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	handler, err := buildHandler(snapshotPath)
	if err != nil {
		t.Fatalf("buildHandler() error = %v; want nil", err)
	}

	requestBody := `{
		"query": "query { autocomplete(prefix: \"java\") { value score } }"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/query",
		strings.NewReader(requestBody),
	)
	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf(
			"status code = %d; want %d; body = %s",
			responseRecorder.Code,
			http.StatusOK,
			responseRecorder.Body.String(),
		)
	}

	responseBody := responseRecorder.Body.String()

	if !strings.Contains(
		responseBody,
		`"value":"javascript"`,
	) {
		t.Errorf(
			"response body = %s; want javascript suggestion",
			responseBody,
		)
	}
}

func TestBuildHandlerReturnsErrorForMissingSnapshot(t *testing.T) {
	snapshotPath := filepath.Join(
		t.TempDir(),
		"missing.json",
	)

	handler, err := buildHandler(snapshotPath)

	if err == nil {
		t.Fatal("buildHandler() error = nil; want error")
	}

	if handler != nil {
		t.Errorf("buildHandler() handler = %#v; want nil", handler)
	}

	if !strings.Contains(err.Error(), "load catalog") {
		t.Errorf(
			"buildHandler() error = %q; want load catalog context",
			err,
		)
	}
}

func TestHealthHandlerReturnsOK(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	responseRecorder := httptest.NewRecorder()

	healthHandler(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf(
			"status code = %d; want %d",
			responseRecorder.Code,
			http.StatusOK,
		)
	}

	if responseRecorder.Body.String() != "ok\n" {
		t.Errorf(
			"response body = %q; want %q",
			responseRecorder.Body.String(),
			"ok\n",
		)
	}
}
