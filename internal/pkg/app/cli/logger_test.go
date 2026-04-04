package cli

import (
	"errors"
	"regexp"
	"strings"
	"testing"
)

func TestSimpleLoggerInfoPrintsLevelAndMessage(t *testing.T) {
	logger := NewSimpleLogger()

	output := captureStdout(t, func() {
		logger.Info("weather loaded")
	})

	pattern := `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] INFO: weather loaded\n$`
	if !regexp.MustCompile(pattern).MatchString(output) {
		t.Fatalf("unexpected log format: %q", output)
	}
}

func TestSimpleLoggerDebugPrintsLevelAndMessage(t *testing.T) {
	logger := NewSimpleLogger()

	output := captureStdout(t, func() {
		logger.Debug("request prepared")
	})

	pattern := `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] DEBUG: request prepared\n$`
	if !regexp.MustCompile(pattern).MatchString(output) {
		t.Fatalf("unexpected log format: %q", output)
	}
}

func TestSimpleLoggerErrorPrintsErrorTextWhenErrorProvided(t *testing.T) {
	logger := NewSimpleLogger()

	output := captureStdout(t, func() {
		logger.Error("request failed", errors.New("boom"))
	})

	pattern := `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] ERROR: request failed: boom\n$`
	if !regexp.MustCompile(pattern).MatchString(output) {
		t.Fatalf("unexpected log format: %q", output)
	}
}

func TestSimpleLoggerErrorPrintsOnlyMessageWhenErrorIsNil(t *testing.T) {
	logger := NewSimpleLogger()

	output := captureStdout(t, func() {
		logger.Error("request failed", nil)
	})

	if !strings.Contains(output, "ERROR: request failed\n") {
		t.Fatalf("expected message without error text, got %q", output)
	}

	if strings.Contains(output, ": <nil>") {
		t.Fatalf("logger should not print nil error, got %q", output)
	}
}
