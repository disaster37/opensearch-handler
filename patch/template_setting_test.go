package patch

import (
	"encoding/json"
	"testing"

	"github.com/disaster37/opensearch/v2"
	"github.com/stretchr/testify/assert"
)

func TestConvertComponentTemplateSetting(t *testing.T) {

	actual := &IndicesGetComponentTemplate{
		IndicesGetComponentTemplate: opensearch.IndicesGetComponentTemplate{
			Template: &opensearch.IndicesGetComponentTemplateData{
				Settings: map[string]any{
					"test": "plop",
					"property": map[string]any{
						"plop": 100,
					},
					"list": []any{
						200,
						300,
					},
				},
			},
		},
	}

	expected := &IndicesGetComponentTemplate{
		IndicesGetComponentTemplate: opensearch.IndicesGetComponentTemplate{
			Template: &opensearch.IndicesGetComponentTemplateData{
				Settings: map[string]any{
					"test": "plop",
					"property": map[string]any{
						"plop": "100",
					},
					"list": []any{
						"200",
						"300",
					},
				},
			},
		},
	}

	acualByte, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	expectedByte, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	acualByte, expectedByte, err = ConvertComponentTemplateSetting(acualByte, expectedByte)
	assert.NoError(t, err)
	assert.Equal(t, expectedByte, acualByte)

}

func TestConvertIndexTemplateSetting(t *testing.T) {

	actual := &IndicesGetIndexTemplate{
		IndicesGetIndexTemplate: opensearch.IndicesGetIndexTemplate{
			Template: &opensearch.IndicesGetIndexTemplateData{
				Settings: map[string]any{
					"test": "plop",
					"property": map[string]any{
						"plop": 100,
					},
					"list": []any{
						200,
						300,
					},
				},
			},
		},
	}

	expected := &IndicesGetIndexTemplate{
		IndicesGetIndexTemplate: opensearch.IndicesGetIndexTemplate{
			Template: &opensearch.IndicesGetIndexTemplateData{
				Settings: map[string]any{
					"test": "plop",
					"property": map[string]any{
						"plop": "100",
					},
					"list": []any{
						"200",
						"300",
					},
				},
			},
		},
	}

	acualByte, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	expectedByte, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	acualByte, expectedByte, err = ConvertIndexTemplateSetting(acualByte, expectedByte)
	assert.NoError(t, err)
	assert.Equal(t, expectedByte, acualByte)

}
