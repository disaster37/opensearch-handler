package samples

import (
	"log"

	"github.com/disaster37/opensearch/v3"
)

func ManageCrossClusterReplication() {
	var (
		ccrRuleStatus           *opensearch.CcrStatusRuleResponse
		expectedCcrRule         *opensearch.CcrRule
		originalCcrRule         *opensearch.CcrRule
		err                     error
		ccrAutoFollowRuleStatus *opensearch.CcrAutoFollowStatus
		expectedAutoFollowRule  *opensearch.CcrAutoFollowRule
		originalAutoFollowRule  *opensearch.CcrAutoFollowRule
	)

	client := GetClient()

	// Create Ccr rule
	ccrRuleCreate := &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	if err = client.CrossClusterReplicationStart("follower-01", ccrRuleCreate); err != nil {
		log.Fatalf("Error when create ccr: %s", err.Error())
	}

	// Get ccr status
	ccrRuleStatus, err = client.CrossClusterReplicationStatus("follower-01")
	if err != nil {
		log.Fatalf("Error when get ccr: %s", err.Error())
	}
	log.Printf("Get ccr successfully: %s\n", ccrRuleStatus.Status)

	// Diff ccr on 3 way merge pattern
	// You need to track somewhere the original ccr.
	// You need to store them after create or update it
	originalCcrRule = &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	expectedCcrRule = &opensearch.CcrRule{
		LeaderAlias: "my-connection-name",
		LeaderIndex: "leader-01",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	diff, err := client.CrossClusterReplicationDiff(ccrRuleCreate, expectedCcrRule, originalCcrRule)
	if err != nil {
		log.Fatalf("Error when diff ccr: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())

		if err = client.CrossClusterReplicationStop("follower-01"); err != nil {
			log.Fatalf("Error when stop ccr: %s", err.Error())
		}

		if err = client.CrossClusterReplicationStart("follower-01", diff.Patched.(*opensearch.CcrRule)); err != nil {
			log.Fatalf("Error when start ccr: %s", err.Error())
		}
	}

	// Delete ccr
	if err = client.CrossClusterReplicationStop("follower-01"); err != nil {
		log.Fatalf("Error when stop ccr: %s", err.Error())
	}

	// Create auto follow rule
	ccrAutoFollowCreate := &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	if err = client.CrossClusterReplicationAutoFollowCreate(ccrAutoFollowCreate); err != nil {
		log.Fatalf("Error when create auto follow: %s", err.Error())
	}

	// Get auto follow status
	ccrAutoFollowRuleStatus, err = client.CrossClusterReplicationAutoFollowStatus("leader-rule")
	if err != nil {
		log.Fatalf("Error when get auto follow: %s", err.Error())
	}
	log.Printf("Get auto follow successfully: %d\n", ccrAutoFollowRuleStatus.NumSuccessStartReplications)

	// Diff auto follow on 3 way merge pattern
	// You need to track somewhere the original auto follow.
	// You need to store them after create or update it
	originalAutoFollowRule = &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	expectedAutoFollowRule = &opensearch.CcrAutoFollowRule{
		LeaderAlias: "my-connection-name",
		Name:        "leader-rule",
		Pattern:     "leader-*",
		UseRoles: opensearch.CcrRuleUseRoles{
			LeaderClusterRole:   "all_access",
			FollowerClusterRole: "all_access",
		},
	}

	diff, err = client.CrossClusterReplicationAutoFollowDiff(ccrAutoFollowCreate, expectedAutoFollowRule, originalAutoFollowRule)
	if err != nil {
		log.Fatalf("Error when diff auto follow: %s", err.Error())
	}
	if !diff.IsEmpty() {
		log.Printf("Found diff %s, you need to update it", diff.String())

		if err = client.CrossClusterReplicationAutoFollowDelete(expectedAutoFollowRule.Name, expectedAutoFollowRule.LeaderAlias); err != nil {
			log.Fatalf("Error when delete auto follow: %s", err.Error())
		}

		if err = client.CrossClusterReplicationAutoFollowCreate(diff.Patched.(*opensearch.CcrAutoFollowRule)); err != nil {
			log.Fatalf("Error create auto follow: %s", err.Error())
		}
	}

	// Delete auto follow
	if err = client.CrossClusterReplicationAutoFollowDelete(ccrAutoFollowCreate.Name, ccrAutoFollowCreate.LeaderAlias); err != nil {
		log.Fatalf("Error when delete auto follow: %s", err.Error())
	}
}
