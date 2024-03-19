package samples

import (
	"log"

	"github.com/disaster37/opensearch/v2"
	"k8s.io/utils/ptr"
)

func ManageUser() {

	var (
		user         *opensearch.SecurityUser
		userToCreate *opensearch.SecurityPutUser
		expectedUser *opensearch.SecurityPutUser
		originalUser *opensearch.SecurityPutUser
		err          error
	)

	client := GetClient()

	// Create user
	userToCreate = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
			Description:   ptr.To[string]("test"),
		},
		Password: ptr.To[string]("my strong password"),
	}

	if err = client.UserUpdate("test", userToCreate); err != nil {
		log.Fatalf("Error when create user: %s", err.Error())
	}

	// Get user
	user, err = client.UserGet("test")
	if err != nil {
		log.Fatalf("Error when get user: %s", err.Error())
	}
	log.Printf("Get user with description %s\n", *user.Description)

	// Diff user on 3 way merge pattern
	// You need to track somewhere the original user.
	// You need to store them after create or update it
	originalUser = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"kibana_user"},
			Description:   ptr.To[string]("test"),
		},
		Password: ptr.To[string]("my strong password"),
	}

	expectedUser = &opensearch.SecurityPutUser{
		SecurityUserBase: opensearch.SecurityUserBase{
			SecurityRoles: []string{"team-a"},
			Description:   ptr.To[string]("test"),
		},
	}

	diff, err := client.UserDiff(userToCreate, expectedUser, originalUser)
	if err != nil {
		log.Fatalf("Error when diff user: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())
		// Update user from diff
		if err = client.UserUpdate("test", diff.Patched.(*opensearch.SecurityPutUser)); err != nil {
			log.Fatalf("Error when update user: %s", err.Error())
		}
	}

	// Delete user
	if err = client.UserDelete("test"); err != nil {
		log.Fatalf("Error when delete user: %s", err.Error())
	}

}
