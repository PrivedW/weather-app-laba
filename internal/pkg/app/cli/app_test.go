package cli

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type weatherInfoStub struct {
	current Current
	lat     float64
	long    float64
	calls   int
}

func (w *weatherInfoStub) GetTemperature(lat, long float64) Current {
	w.calls++
	w.lat = lat
	w.long = long
	return w.current
}

type loggerStub struct{}

func (l *loggerStub) Info(string)           {}
func (l *loggerStub) Debug(string)          {}
func (l *loggerStub) Error(string, error)   {}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe stdout: %v", err)
	}

	os.Stdout = writer
	fn()
	_ = writer.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		t.Fatalf("read stdout: %v", err)
	}

	return buf.String()
}

func TestCLIAppRunPrintsTemperatureForConfiguredCoordinates(t *testing.T) {
	wi := &weatherInfoStub{
		current: Current{Temp: 21.35},
	}
	app := New(&loggerStub{}, wi)

	output := captureStdout(t, func() {
		if err := app.Run(); err != nil {
			t.Fatalf("Run() returned error: %v", err)
		}
	})

	expected := "Температура воздуха- 21.35 градусов цельсия\n"
	if output != expected {
		t.Fatalf("unexpected output:\nwant: %q\ngot:  %q", expected, output)
	}

	if wi.calls != 1 {
		t.Fatalf("expected weather info call count to be 1, got %d", wi.calls)
	}

	if wi.lat != 53.6688 || wi.long != 23.8223 {
		t.Fatalf("unexpected coordinates: lat=%f long=%f", wi.lat, wi.long)
	}
}
