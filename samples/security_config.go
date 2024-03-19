package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageSecurityConfig() {

	var (
		securityConfig         *opensearch.SecurityGetConfigResponse
		expectedSecurityConfig *opensearch.SecurityConfig
		originalSecurityConfig *opensearch.SecurityConfig
		err                    error
	)

	client := GetClient()

	// Update security config
	securityConfig = &opensearch.SecurityGetConfigResponse{
		Config: opensearch.SecurityConfig{
			Dynamic: opensearch.SecurityConfigDynamic{
				DoNotFailOnForbidden:      ptr.To[bool](true),
				DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
				Http: &opensearch.SecurityConfigHttp{
					AnonymousAuthEnabled: ptr.To[bool](false),
				},
			},
		},
	}

	if err = client.SecurityConfigUpdate(&securityConfig.Config); err != nil {
		log.Fatalf("Error when update security config: %s", err.Error())
	}

	// Get security config
	securityConfig, err = client.SecurityConfigGet()
	if err != nil {
		log.Fatalf("Error when get security config: %s", err.Error())
	}
	log.Print("Get security config successfully\n")

	// Diff security config on 3 way merge pattern
	// You need to track somewhere the original security config.
	// You need to store them after create or update it
	originalSecurityConfig = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
			Http: &opensearch.SecurityConfigHttp{
				AnonymousAuthEnabled: ptr.To[bool](false),
			},
		},
	}

	expectedSecurityConfig = &opensearch.SecurityConfig{
		Dynamic: opensearch.SecurityConfigDynamic{
			DoNotFailOnForbidden:      ptr.To[bool](true),
			DoNotFailOnForbiddenEmpty: ptr.To[bool](true),
			Http: &opensearch.SecurityConfigHttp{
				AnonymousAuthEnabled: ptr.To[bool](true),
			},
		},
	}

	diff, err := client.SecurityConfigDiff(&securityConfig.Config, expectedSecurityConfig, originalSecurityConfig)
	if err != nil {
		log.Fatalf("Error when diff security config: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update security config from diff
		if err = client.SecurityConfigUpdate(diff.Patched.(*opensearch.SecurityConfig)); err != nil {
			log.Fatalf("Error when update security config: %s", err.Error())
		}
	}

}
