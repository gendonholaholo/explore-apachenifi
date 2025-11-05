package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ghawsshafadonia/nifiparser/internal/config"
)

func TestNew(t *testing.T) {
	filePath := "/test/path/file.json"
	p := New(filePath)

	if p == nil {
		t.Fatal("Expected non-nil parser")
	}

	if p.filePath != filePath {
		t.Errorf("Expected filePath %s, got %s", filePath, p.filePath)
	}
}

func TestParse_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.json")

	content := `{"flowId":"test-123","flowName":"Test Flow","version":"1.0.0","processors":[],"connections":[]}`

	if err := os.WriteFile(tmpFile, []byte(content), 0o600); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	p := New(tmpFile)
	flow, err := p.Parse()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if flow.FlowID != "test-123" {
		t.Errorf("Expected flowId 'test-123', got '%s'", flow.FlowID)
	}

	if flow.FlowName != "Test Flow" {
		t.Errorf("Expected flowName 'Test Flow', got '%s'", flow.FlowName)
	}
}

func TestParse_FileNotFound(t *testing.T) {
	p := New("/nonexistent/file.json")
	_, err := p.Parse()

	if err == nil {
		t.Fatal("Expected error for nonexistent file, got nil")
	}
}

func TestParse_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "invalid.json")

	content := `{invalid json}`

	if err := os.WriteFile(tmpFile, []byte(content), 0o600); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	p := New(tmpFile)
	_, err := p.Parse()

	if err == nil {
		t.Fatal("Expected error for invalid JSON, got nil")
	}
}

func TestValidateFlow_ValidFlow(t *testing.T) {
	flow := &config.NiFiFlow{
		FlowID:   "test-id",
		FlowName: "Test Name",
	}

	err := ValidateFlow(flow)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestValidateFlow_MissingFlowID(t *testing.T) {
	flow := &config.NiFiFlow{
		FlowName: "Test Name",
	}

	err := ValidateFlow(flow)
	if err == nil {
		t.Fatal("Expected error for missing flowID, got nil")
	}
}

func TestValidateFlow_MissingFlowName(t *testing.T) {
	flow := &config.NiFiFlow{
		FlowID: "test-id",
	}

	err := ValidateFlow(flow)
	if err == nil {
		t.Fatal("Expected error for missing flowName, got nil")
	}
}
