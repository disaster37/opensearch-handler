package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageTenant() {

	var (
		tenant         *opensearch.SecurityTenant
		expectedTenant *opensearch.SecurityPutTenant
		originalTenant *opensearch.SecurityPutTenant
		err            error
	)

	client := GetClient()

	// Create tenant
	tenant = &opensearch.SecurityTenant{
		SecurityPutTenant: opensearch.SecurityPutTenant{
			Description: ptr.To[string]("test"),
		},
	}

	if err = client.TenantUpdate("test", &tenant.SecurityPutTenant); err != nil {
		log.Fatalf("Error when create tenant: %s", err.Error())
	}

	// Get tenant
	tenant, err = client.TenantGet("test")
	if err != nil {
		log.Fatalf("Error when get tenant: %s", err.Error())
	}
	log.Printf("Get tenant with description %s\n", *tenant.Description)

	// Diff tenant on 3 way merge pattern
	// You need to track somewhere the original tenant.
	// You need to store them after create or update it
	originalTenant = &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("test"),
	}

	expectedTenant = &opensearch.SecurityPutTenant{
		Description: ptr.To[string]("my new description"),
	}

	diff, err := client.TenantDiff(&tenant.SecurityPutTenant, expectedTenant, originalTenant)
	if err != nil {
		log.Fatalf("Error when diff tenant: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update tenant from diff
		if err = client.TenantUpdate("test", diff.Patched.(*opensearch.SecurityPutTenant)); err != nil {
			log.Fatalf("Error when update tenant: %s", err.Error())
		}
	}

	// Delete tenant
	if err = client.TenantDelete("test"); err != nil {
		log.Fatalf("Error when delete tenant: %s", err.Error())
	}

}
