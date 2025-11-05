// Package config provides data models for NiFi configuration
package config

// NiFiFlow represents the top-level NiFi flow configuration
type NiFiFlow struct {
	FlowID      string       `json:"flowId"`
	FlowName    string       `json:"flowName"`
	Version     string       `json:"version"`
	Processors  []Processor  `json:"processors"`
	Connections []Connection `json:"connections"`
}

// Processor represents a NiFi processor component
type Processor struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
	Status     string            `json:"status"`
}

// Connection represents a connection between processors
type Connection struct {
	ID           string `json:"id"`
	Source       string `json:"source"`
	Destination  string `json:"destination"`
	Relationship string `json:"relationship"`
}
