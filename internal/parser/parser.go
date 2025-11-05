// Package parser handles reading and parsing NiFi configuration files
package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghawsshafadonia/nifiparser/internal/config"
)

// Parser handles NiFi configuration parsing operations
type Parser struct {
	filePath string
}

// New creates a new Parser instance
func New(filePath string) *Parser {
	return &Parser{
		filePath: filePath,
	}
}

// Parse reads and parses the NiFi configuration file
func (p *Parser) Parse() (*config.NiFiFlow, error) {
	// Read file
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse JSON
	var flow config.NiFiFlow
	if err := json.Unmarshal(data, &flow); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &flow, nil
}

// ValidateFlow performs basic validation on the parsed flow
func ValidateFlow(flow *config.NiFiFlow) error {
	if flow.FlowID == "" {
		return fmt.Errorf("flow ID is required")
	}
	if flow.FlowName == "" {
		return fmt.Errorf("flow name is required")
	}
	return nil
}
