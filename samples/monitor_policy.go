package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
)

func ManageMonitor() {

	var (
		monitor         *opensearch.AlertingGetMonitorResponse
		expectedMonitor *opensearch.AlertingMonitor
		originalMonitor *opensearch.AlertingMonitor
		err             error
	)

	client := GetClient()

	// Create monitor
	monitor = &opensearch.AlertingGetMonitorResponse{
		Monitor: opensearch.AlertingGetMonitor{
			AlertingMonitor: opensearch.AlertingMonitor{
				Type: "monitor",
				Name: "test",
			},
		},
	}

	if err = client.MonitorCreate(&monitor.Monitor.AlertingMonitor); err != nil {
		log.Fatalf("Error when create monitor: %s", err.Error())
	}

	// Get monitor
	monitor, err = client.MonitorGet("test")
	if err != nil {
		log.Fatalf("Error when get monitor: %s", err.Error())
	}
	log.Print("Get monitor successfully\n")

	// Diff monitor on 3 way merge pattern
	// You need to track somewhere the original monitor.
	// You need to store them after create or update it
	originalMonitor = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
	}

	expectedMonitor = &opensearch.AlertingMonitor{
		Type: "monitor",
		Name: "test",
		Schedule: map[string]any{
			"period": map[string]any{
				"interval": 1,
				"unit":     "MINUTES",
			},
		},
	}

	diff, err := client.MonitorDiff(&monitor.Monitor.AlertingMonitor, expectedMonitor, originalMonitor)
	if err != nil {
		log.Fatalf("Error when diff monitor: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update monitor from diff
		if err = client.MonitorUpdate("test", monitor.SequenceNumber, monitor.PrimaryTerm, diff.Patched.(*opensearch.AlertingMonitor)); err != nil {
			log.Fatalf("Error when update monitor: %s", err.Error())
		}
	}

	// Delete monitor
	if err = client.MonitorDelete("test"); err != nil {
		log.Fatalf("Error when delete monitor: %s", err.Error())
	}

}
