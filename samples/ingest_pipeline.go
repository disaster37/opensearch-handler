package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
)

func ManageIngestPipeline() {

	var (
		ingestPipeline         *opensearch.IngestGetPipeline
		expectedIngestPipeline *opensearch.IngestGetPipeline
		originalIngestPipeline *opensearch.IngestGetPipeline
		err                    error
	)

	client := GetClient()

	// Create ingest pipeline
	ingestPipeline = &opensearch.IngestGetPipeline{
		Description: "test",
		Processors: []map[string]interface{}{
			{
				"uppercase": map[string]any{
					"field": "name",
				},
			},
		},
	}

	if err = client.IngestPipelineUpdate("test", ingestPipeline); err != nil {
		log.Fatalf("Error when create ingest pipeline: %s", err.Error())
	}

	// Get ingest pipeline
	ingestPipeline, err = client.IngestPipelineGet("test")
	if err != nil {
		log.Fatalf("Error when get ingest pipeline: %s", err.Error())
	}
	log.Printf("Get ingest pipeline with version %d\n", ingestPipeline.Version)

	// Diff ingest pipeline on 3 way merge pattern
	// You need to track somewhere the original ingest pipeline.
	// You need to store them after create or update it
	originalIngestPipeline = &opensearch.IngestGetPipeline{
		Description: "test",
		Processors: []map[string]interface{}{
			{
				"uppercase": map[string]any{
					"field": "name",
				},
			},
		},
	}

	expectedIngestPipeline = &opensearch.IngestGetPipeline{
		Description: "test",
		Processors: []map[string]interface{}{
			{
				"uppercase": map[string]any{
					"field": "country",
				},
			},
		},
	}

	diff, err := client.IngestPipelineDiff(ingestPipeline, expectedIngestPipeline, originalIngestPipeline)
	if err != nil {
		log.Fatalf("Error when diff ingest pipeline: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update ingest pipeline from diff
		if err = client.IngestPipelineUpdate("test", diff.Patched.(*opensearch.IngestGetPipeline)); err != nil {
			log.Fatalf("Error when update ingest pipeline: %s", err.Error())
		}
	}

	// Delete ingest pipeline
	if err = client.IngestPipelineDelete("test"); err != nil {
		log.Fatalf("Error when delete ingest pipeline: %s", err.Error())
	}

}
