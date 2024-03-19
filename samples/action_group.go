package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageActionGroup() {

	var (
		actionGroup         *opensearch.SecurityActionGroup
		expectedActionGroup *opensearch.SecurityPutActionGroup
		originalActionGroup *opensearch.SecurityPutActionGroup
		err                 error
	)

	client := GetClient()

	// Create action group
	actionGroup = &opensearch.SecurityActionGroup{
		SecurityPutActionGroup: opensearch.SecurityPutActionGroup{
			AllowedActions: []string{"cluster_all"},
			Description:    ptr.To[string]("test"),
		},
	}

	if err = client.ActionGroupUpdate("test", &actionGroup.SecurityPutActionGroup); err != nil {
		log.Fatalf("Error when create action group: %s", err.Error())
	}

	// Get action group
	actionGroup, err = client.ActionGroupGet("test")
	if err != nil {
		log.Fatalf("Error when get action group: %s", err.Error())
	}
	log.Printf("Get action group with description %s\n", *actionGroup.Description)

	// Diff action group on 3 way merge pattern
	// You need to track somewhere the original action group.
	// You need to store them after create or update it
	originalActionGroup = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"cluster_all"},
		Description:    ptr.To[string]("test"),
	}

	expectedActionGroup = &opensearch.SecurityPutActionGroup{
		AllowedActions: []string{"crud", "search"},
		Description:    ptr.To[string]("test"),
	}

	diff, err := client.ActionGroupDiff(&actionGroup.SecurityPutActionGroup, expectedActionGroup, originalActionGroup)
	if err != nil {
		log.Fatalf("Error when diff action group: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update action group from diff
		if err = client.ActionGroupUpdate("test", diff.Patched.(*opensearch.SecurityPutActionGroup)); err != nil {
			log.Fatalf("Error when update action group: %s", err.Error())
		}
	}

	// Delete actionGroup
	if err = client.ActionGroupDelete("test"); err != nil {
		log.Fatalf("Error when delete action group: %s", err.Error())
	}

}
