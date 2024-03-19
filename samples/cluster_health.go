package samples

import "log"

func GetClusterHealth() {

	client := GetClient()

	health, err := client.ClusterHealth()
	if err != nil {
		log.Fatalf("Error when get cluster health: %s", err.Error())
	}
	log.Printf("The cluster health is %s", health.Status)

}
