package samples

import (
	"log"
	"strings"

	"github.com/disaster37/opensearch/v2"
)

func ManageRoleMapping() {

	var (
		roleMapping         *opensearch.SecurityRoleMapping
		expectedRoleMapping *opensearch.SecurityPutRoleMapping
		originalRoleMapping *opensearch.SecurityPutRoleMapping
		err                 error
	)

	client := GetClient()

	// Create role mapping
	roleMapping = &opensearch.SecurityRoleMapping{
		SecurityPutRoleMapping: opensearch.SecurityPutRoleMapping{
			BackendRoles: []string{"ad_group"},
		},
	}

	if err = client.RoleMappingUpdate("test", &roleMapping.SecurityPutRoleMapping); err != nil {
		log.Fatalf("Error when create role mapping: %s", err.Error())
	}

	// Get role mapping
	roleMapping, err = client.RoleMappingGet("test")
	if err != nil {
		log.Fatalf("Error when get role mapping: %s", err.Error())
	}
	log.Printf("Get role mapping with backend roles %s\n", strings.Join(roleMapping.BackendRoles, ","))

	// Diff role mapping on 3 way merge pattern
	// You need to track somewhere the original role mapping.
	// You need to store them after create or update it
	originalRoleMapping = &opensearch.SecurityPutRoleMapping{
		BackendRoles: []string{"ad_group"},
	}

	expectedRoleMapping = &opensearch.SecurityPutRoleMapping{
		BackendRoles: []string{"ad_group_new_team"},
	}

	diff, err := client.RoleMappingDiff(&roleMapping.SecurityPutRoleMapping, expectedRoleMapping, originalRoleMapping)
	if err != nil {
		log.Fatalf("Error when diff role mapping: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update role mapping from diff
		if err = client.RoleMappingUpdate("test", diff.Patched.(*opensearch.SecurityPutRoleMapping)); err != nil {
			log.Fatalf("Error when update role mapping: %s", err.Error())
		}
	}

	// Delete role mapping
	if err = client.RoleMappingDelete("test"); err != nil {
		log.Fatalf("Error when delete role mapping: %s", err.Error())
	}

}
