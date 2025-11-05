// Package display handles output formatting and display logic
package display

import (
	"fmt"
	"strings"

	"github.com/ghawsshafadonia/nifiparser/internal/config"
)

// Displayer handles displaying NiFi flow information
type Displayer struct {
	flow *config.NiFiFlow
}

// New creates a new Displayer instance
func New(flow *config.NiFiFlow) *Displayer {
	return &Displayer{
		flow: flow,
	}
}

// PrintSummary displays a summary of the NiFi flow
func (d *Displayer) PrintSummary() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("NiFi Flow Summary: %s\n", d.flow.FlowName)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Flow ID:       %s\n", d.flow.FlowID)
	fmt.Printf("Version:       %s\n", d.flow.Version)
	fmt.Printf("Processors:    %d\n", len(d.flow.Processors))
	fmt.Printf("Connections:   %d\n", len(d.flow.Connections))
	fmt.Println()
}

// PrintProcessors displays detailed information about processors
func (d *Displayer) PrintProcessors() {
	if len(d.flow.Processors) == 0 {
		fmt.Println("No processors found.")
		return
	}

	fmt.Println("Processors:")
	fmt.Println(strings.Repeat("-", 60))
	for i, proc := range d.flow.Processors {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, proc.Name, proc.ID)
		fmt.Printf("   Type:   %s\n", proc.Type)
		fmt.Printf("   Status: %s\n", proc.Status)
		if len(proc.Properties) > 0 {
			fmt.Println("   Properties:")
			for key, value := range proc.Properties {
				fmt.Printf("     - %s: %s\n", key, value)
			}
		}
		fmt.Println()
	}
}

// PrintConnections displays connection information
func (d *Displayer) PrintConnections() {
	if len(d.flow.Connections) == 0 {
		fmt.Println("No connections found.")
		return
	}

	fmt.Println("Connections:")
	fmt.Println(strings.Repeat("-", 60))
	for i, conn := range d.flow.Connections {
		fmt.Printf("%d. %s -> %s (%s)\n",
			i+1, conn.Source, conn.Destination, conn.Relationship)
	}
	fmt.Println()
}

// PrintStatistics displays flow statistics
func (d *Displayer) PrintStatistics() {
	runningCount := 0
	stoppedCount := 0

	for _, proc := range d.flow.Processors {
		switch proc.Status {
		case "running":
			runningCount++
		case "stopped":
			stoppedCount++
		}
	}

	fmt.Println("Statistics:")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Running Processors:  %d\n", runningCount)
	fmt.Printf("Stopped Processors:  %d\n", stoppedCount)
	fmt.Printf("Total Connections:   %d\n", len(d.flow.Connections))
	fmt.Println()
}
