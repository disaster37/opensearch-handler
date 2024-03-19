package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageAuditConfig() {

	var (
		auditConfig         *opensearch.SecurityGetAuditResponse
		expectedAuditConfig *opensearch.SecurityAudit
		originalAuditConfig *opensearch.SecurityAudit
		err                 error
	)

	client := GetClient()

	// Create audit config
	auditConfig = &opensearch.SecurityGetAuditResponse{
		Config: opensearch.SecurityAudit{
			Enabled: ptr.To[bool](true),
			Audit: opensearch.SecurityAuditSpec{
				IgnoreUsers: []string{"prometheus"},
			},
		},
	}

	if err = client.SecurityAuditUpdate(&auditConfig.Config); err != nil {
		log.Fatalf("Error when create audit config: %s", err.Error())
	}

	// Get audit config
	auditConfig, err = client.SecurityAuditGet()
	if err != nil {
		log.Fatalf("Error when get audit config: %s", err.Error())
	}
	log.Print("Get audit config successfully\n")

	// Diff audit config on 3 way merge pattern
	// You need to track somewhere the original audit config.
	// You need to store them after create or update it
	originalAuditConfig = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"prometheus"},
		},
	}

	expectedAuditConfig = &opensearch.SecurityAudit{
		Enabled: ptr.To[bool](true),
		Audit: opensearch.SecurityAuditSpec{
			IgnoreUsers: []string{"prometheus", "monitoring"},
		},
	}

	diff, err := client.SecurityAuditDiff(&auditConfig.Config, expectedAuditConfig, originalAuditConfig)
	if err != nil {
		log.Fatalf("Error when diff audit config: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update audit config from diff
		if err = client.SecurityAuditUpdate(diff.Patched.(*opensearch.SecurityAudit)); err != nil {
			log.Fatalf("Error when update audit config: %s", err.Error())
		}
	}

}
