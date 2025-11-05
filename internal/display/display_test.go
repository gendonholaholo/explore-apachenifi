package display

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ghawsshafadonia/nifiparser/internal/config"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNew(t *testing.T) {
	flow := &config.NiFiFlow{
		FlowID:   "test-123",
		FlowName: "Test Flow",
	}

	d := New(flow)

	if d == nil {
		t.Fatal("Expected non-nil displayer")
	}

	if d.flow != flow {
		t.Error("Expected flow to match")
	}
}

func TestPrintSummary(t *testing.T) {
	flow := &config.NiFiFlow{
		FlowID:      "test-123",
		FlowName:    "Test Flow",
		Version:     "1.0.0",
		Processors:  []config.Processor{{ID: "p1"}},
		Connections: []config.Connection{{ID: "c1"}},
	}

	d := New(flow)
	output := captureOutput(func() {
		d.PrintSummary()
	})

	if !strings.Contains(output, "Test Flow") {
		t.Error("Expected output to contain flow name")
	}

	if !strings.Contains(output, "test-123") {
		t.Error("Expected output to contain flow ID")
	}
}

func TestPrintProcessors_NoProcessors(t *testing.T) {
	flow := &config.NiFiFlow{
		Processors: []config.Processor{},
	}

	d := New(flow)
	output := captureOutput(func() {
		d.PrintProcessors()
	})

	if !strings.Contains(output, "No processors found") {
		t.Error("Expected message about no processors")
	}
}

func TestPrintStatistics(t *testing.T) {
	flow := &config.NiFiFlow{
		Processors: []config.Processor{
			{Status: "running"},
			{Status: "running"},
			{Status: "stopped"},
		},
		Connections: []config.Connection{{}, {}},
	}

	d := New(flow)
	output := captureOutput(func() {
		d.PrintStatistics()
	})

	if !strings.Contains(output, "Running Processors:  2") {
		t.Error("Expected 2 running processors in output")
	}

	if !strings.Contains(output, "Stopped Processors:  1") {
		t.Error("Expected 1 stopped processor in output")
	}
}
