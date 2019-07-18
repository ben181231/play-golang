package leak

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"go.uber.org/goleak"
)

func TestGetSumHandler(t *testing.T) {
	// -- Given --
	goroutineCount := runtime.NumGoroutine()
	handler := GetSumHandler()
	req := httptest.NewRequest(
		http.MethodGet,
		"http://somewhere.else",
		strings.NewReader(
			`
				{
					"first": 1,
					"second": 1,
					"third": 1
				}
			`,
		),
	)
	recorder := httptest.NewRecorder()

	// -- When --
	handler.ServeHTTP(recorder, req)

	// -- Then --
	statusCode := recorder.Code
	if statusCode != 200 {
		t.Errorf("Expected get status 200, but got %d", statusCode)
	}
	goroutineCountAfter := runtime.NumGoroutine()
	if goroutineCountAfter != goroutineCount {
		t.Errorf(
			"Expected %d goroutine running, but found %d",
			goroutineCount,
			goroutineCountAfter,
		)
	}
}

func TestGetSumHandler_WithGoLeak(t *testing.T) {
	defer goleak.VerifyNoLeaks(t)

	// -- Given --
	handler := GetSumHandler()
	req := httptest.NewRequest(
		http.MethodGet,
		"http://somewhere.else",
		strings.NewReader(
			`
				{
					"first": 1,
					"second": 1,
					"third": 1
				}
			`,
		),
	)
	recorder := httptest.NewRecorder()

	// -- When --
	handler.ServeHTTP(recorder, req)

	// -- Then --
	statusCode := recorder.Code
	if statusCode != 200 {
		t.Errorf("Expected get status 200, but got %d", statusCode)
	}
}
