package weather

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
)

type loggerSpy struct {
	debugs []string
	errors []string
}

func (l *loggerSpy) Info(string) {}

func (l *loggerSpy) Debug(msg string) {
	l.debugs = append(l.debugs, msg)
}

func (l *loggerSpy) Error(msg string, err error) {
	if err != nil {
		l.errors = append(l.errors, msg+": "+err.Error())
		return
	}

	l.errors = append(l.errors, msg)
}

func TestGetTemperatureLoadsWeatherOnceAndCachesValue(t *testing.T) {
	logger := &loggerSpy{}
	oldHTTPGet := httpGet
	t.Cleanup(func() {
		httpGet = oldHTTPGet
	})

	var calls int
	httpGet = func(gotURL string) (*http.Response, error) {
		calls++
		expectedURL := "https://api.open-meteo.com/v1/forecast?latitude=53.100000&longitude=23.200000&current=temperature_2m"
		if gotURL != expectedURL {
			t.Fatalf("unexpected URL:\nwant: %s\ngot:  %s", expectedURL, gotURL)
		}

		body := io.NopCloser(strings.NewReader(`{"current":{"temperature_2m":17.5}}`))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       body,
		}, nil
	}

	wi := New(logger)

	first := wi.GetTemperature(53.1, 23.2)
	second := wi.GetTemperature(53.1, 23.2)

	if first.Temp != 17.5 {
		t.Fatalf("expected first temperature to be 17.5, got %.2f", first.Temp)
	}

	if second.Temp != 17.5 {
		t.Fatalf("expected cached temperature to be 17.5, got %.2f", second.Temp)
	}

	if calls != 1 {
		t.Fatalf("expected one HTTP call, got %d", calls)
	}

	if len(logger.debugs) != 2 {
		t.Fatalf("expected two debug logs, got %d", len(logger.debugs))
	}
}

func TestGetTemperatureReturnsZeroValueWhenRequestFails(t *testing.T) {
	logger := &loggerSpy{}
	oldHTTPGet := httpGet
	t.Cleanup(func() {
		httpGet = oldHTTPGet
	})

	httpGet = func(string) (*http.Response, error) {
		return nil, errors.New("network down")
	}

	wi := New(logger)
	current := wi.GetTemperature(53.1, 23.2)

	if current != (cli.Current{}) {
		t.Fatalf("expected zero value current, got %+v", current)
	}

	if len(logger.errors) < 2 {
		t.Fatalf("expected error logs to be written, got %d", len(logger.errors))
	}

	if !strings.Contains(logger.errors[0], "can`t get weather data") {
		t.Fatalf("unexpected first error log: %q", logger.errors[0])
	}

	if !strings.Contains(logger.errors[len(logger.errors)-1], "can`t load weather info") {
		t.Fatalf("unexpected final error log: %q", logger.errors[len(logger.errors)-1])
	}
}

func TestGetTemperatureReturnsZeroValueWhenJSONIsInvalid(t *testing.T) {
	logger := &loggerSpy{}
	oldHTTPGet := httpGet
	t.Cleanup(func() {
		httpGet = oldHTTPGet
	})

	httpGet = func(string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"current":`)),
		}, nil
	}

	wi := New(logger)
	current := wi.GetTemperature(53.1, 23.2)

	if current != (cli.Current{}) {
		t.Fatalf("expected zero value current, got %+v", current)
	}

	if len(logger.errors) == 0 {
		t.Fatal("expected unmarshal error to be logged")
	}

	found := false
	for _, msg := range logger.errors {
		if strings.Contains(msg, "can`t unmarshal json data") {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected unmarshal log, got %v", logger.errors)
	}
}
