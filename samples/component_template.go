package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
)

func ManageComponentTemplate() {
	var (
		componentTemplate         *opensearch.IndicesGetComponentTemplate
		expectedComponentTemplate *opensearch.IndicesGetComponentTemplate
		originalComponentTemplate *opensearch.IndicesGetComponentTemplate
		err                       error
	)

	client := GetClient()

	// Create component template
	componentTemplate = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	if err = client.ComponentTemplateUpdate("test", componentTemplate); err != nil {
		log.Fatalf("Error when create component template: %s", err.Error())
	}

	// Get component template
	componentTemplate, err = client.ComponentTemplateGet("test")
	if err != nil {
		log.Fatalf("Error when get component template: %s", err.Error())
	}
	log.Printf("Get component template with version %d", componentTemplate.Version)

	// Diff component template on 3 way merge pattern
	// You need to track somewhere the original component template.
	// You need to store them after create or update it
	originalComponentTemplate = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "5s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	expectedComponentTemplate = &opensearch.IndicesGetComponentTemplate{
		Template: &opensearch.IndicesGetComponentTemplateData{
			Settings: map[string]any{
				"index.refresh_interval": "10s",
			},
			Mappings: map[string]any{
				"_source.enabled":           true,
				"properties.host_name.type": "keyword",
			},
		},
	}

	diff, err := client.ComponentTemplateDiff(componentTemplate, expectedComponentTemplate, originalComponentTemplate)
	if err != nil {
		log.Fatalf("Error when diff component template: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update component template from diff
		if err = client.ComponentTemplateUpdate("test", diff.Patched.(*opensearch.IndicesGetComponentTemplate)); err != nil {
			log.Fatalf("Error when update component template: %s", err.Error())
		}
	}

	// Delete component template
	if err = client.ComponentTemplateDelete("test"); err != nil {
		log.Fatalf("Error when delete component template: %s", err.Error())
	}
}
