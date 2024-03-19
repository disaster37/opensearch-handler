package samples

import (
	"crypto/tls"
	"log"
	"net/http"

	opensearchhandler "github.com/disaster37/opensearch-handler/v2"
	"github.com/disaster37/opensearch/v2/config"
	"github.com/sirupsen/logrus"
	"k8s.io/utils/ptr"
)

func GetClient() opensearchhandler.OpensearchHandler {

	client, err := opensearchhandler.NewOpensearchHandler(&config.Config{
		URLs:        []string{"https://127.0.0.1:9200"},
		Username:    "admin",
		Password:    "admin",
		Sniff:       ptr.To[bool](false),
		Healthcheck: ptr.To[bool](false),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}, logrus.NewEntry(logrus.New()))

	if err != nil {
		log.Fatalf("Error when init Opensearch handler: %s", err.Error())
	}

	return client
}
