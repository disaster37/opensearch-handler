package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
)

func ManageIndexTemplate() {

	var (
		indexTemplate         *opensearch.IndicesGetIndexTemplate
		expectedIndexTemplate *opensearch.IndicesGetIndexTemplate
		originalIndexTemplate *opensearch.IndicesGetIndexTemplate
		err                   error
	)

	client := GetClient()

	// Create index template
	indexTemplate = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	if err = client.IndexTemplateUpdate("test", indexTemplate); err != nil {
		log.Fatalf("Error when create index template: %s", err.Error())
	}

	// Get index template
	indexTemplate, err = client.IndexTemplateGet("test")
	if err != nil {
		log.Fatalf("Error when get index template: %s", err.Error())
	}
	log.Printf("Get index template with version %d\n", indexTemplate.Version)

	// Diff index template on 3 way merge pattern
	// You need to track somewhere the original index template.
	// You need to store them after create or update it
	originalIndexTemplate = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
		},
	}

	expectedIndexTemplate = &opensearch.IndicesGetIndexTemplate{
		IndexPatterns: []string{"test-index-template"},
		Priority:      2,
		Template: &opensearch.IndicesGetIndexTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "10s",
			},
		},
	}

	diff, err := client.IndexTemplateDiff(indexTemplate, expectedIndexTemplate, originalIndexTemplate)
	if err != nil {
		log.Fatalf("Error when diff index template: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update index template from diff
		if err = client.IndexTemplateUpdate("test", diff.Patched.(*opensearch.IndicesGetIndexTemplate)); err != nil {
			log.Fatalf("Error when update index template: %s", err.Error())
		}
	}

	// Delete index template
	if err = client.SnapshotRepositoryDelete("test"); err != nil {
		log.Fatalf("Error when delete index template: %s", err.Error())
	}

}
