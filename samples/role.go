package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageRole() {

	var (
		role         *opensearch.SecurityRole
		expectedRole *opensearch.SecurityPutRole
		originalRole *opensearch.SecurityPutRole
		err          error
	)

	client := GetClient()

	// Create role
	role = &opensearch.SecurityRole{
		SecurityPutRole: opensearch.SecurityPutRole{
			Description:        ptr.To[string]("test"),
			ClusterPermissions: []string{"all"},
			IndexPermissions: []opensearch.SecurityIndexPermissions{
				{
					IndexPatterns:  []string{"logstash-*"},
					AllowedActions: []string{"read"},
				},
			},
		},
	}

	if err = client.RoleUpdate("test", &role.SecurityPutRole); err != nil {
		log.Fatalf("Error when create role: %s", err.Error())
	}

	// Get role
	role, err = client.RoleGet("test")
	if err != nil {
		log.Fatalf("Error when get role: %s", err.Error())
	}
	log.Printf("Get role with description %s\n", *role.Description)

	// Diff role on 3 way merge pattern
	// You need to track somewhere the original role.
	// You need to store them after create or update it
	originalRole = &opensearch.SecurityPutRole{
		Description:        ptr.To[string]("test"),
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	expectedRole = &opensearch.SecurityPutRole{
		Description:        ptr.To[string]("Update the description"),
		ClusterPermissions: []string{"all"},
		IndexPermissions: []opensearch.SecurityIndexPermissions{
			{
				IndexPatterns:  []string{"logstash-*"},
				AllowedActions: []string{"read"},
			},
		},
	}

	diff, err := client.RoleDiff(&role.SecurityPutRole, expectedRole, originalRole)
	if err != nil {
		log.Fatalf("Error when diff role: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update role from diff
		if err = client.RoleUpdate("test", diff.Patched.(*opensearch.SecurityPutRole)); err != nil {
			log.Fatalf("Error when update role: %s", err.Error())
		}
	}

	// Delete role
	if err = client.RoleDelete("test"); err != nil {
		log.Fatalf("Error when delete role: %s", err.Error())
	}

}
