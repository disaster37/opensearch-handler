package patch

import (
	"encoding/json"
	"testing"

	"github.com/disaster37/opensearch/v2"
	"github.com/stretchr/testify/assert"
)

func TestRemoveEnvironmentVariableContendTest(t *testing.T) {

	actual := &opensearch.SecurityGetConfigResponse{
		Config: opensearch.SecurityConfig{
			Dynamic: opensearch.SecurityConfigDynamic{
				Authc: map[string]opensearch.SecurityConfigAuthc{
					"ldap": {
						AuthenticationBackend: &opensearch.SecurityConfigAuthenticationBackend{
							Config: map[string]any{
								"username": "my user",
								"password": "my password",
								"dd":       "toto",
							},
						},
					},
				},
			},
		},
	}

	expected := &opensearch.SecurityGetConfigResponse{
		Config: opensearch.SecurityConfig{
			Dynamic: opensearch.SecurityConfigDynamic{
				Authc: map[string]opensearch.SecurityConfigAuthc{
					"ldap": {
						AuthenticationBackend: &opensearch.SecurityConfigAuthenticationBackend{
							Config: map[string]any{
								"username": "my user",
								"password": "${env.PASSWORD}",
								"dd":       "toto",
							},
						},
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

	acualByte, expectedByte, err = RemoveEnvironmentVariableContend(acualByte, expectedByte)
	assert.NoError(t, err)
	assert.Equal(t, expectedByte, acualByte)

}
