package opensearchhandler

import (
	"net/http"
	"testing"

	"github.com/disaster37/opensearch/v2"
	"github.com/jarcoal/httpmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

const baseURL = "http://localhost:9200"

type OpensearchHandlerTestSuite struct {
	suite.Suite
	opensearchHandler OpensearchHandler
}

func TestOpensearchHandlerSuite(t *testing.T) {
	suite.Run(t, new(OpensearchHandlerTestSuite))
}

func (t *OpensearchHandlerTestSuite) SetupTest() {

	httpClient := http.DefaultClient
	httpClient.Transport = httpmock.DefaultTransport
	httpmock.Activate()

	client, err := opensearch.NewClient(opensearch.SetURL(baseURL), opensearch.SetHttpClient(httpClient), opensearch.SetHealthcheck(false), opensearch.SetSniff(false))
	if err != nil {
		panic(err)
	}

	t.opensearchHandler = &OpensearchHandlerImpl{
		client: client,
		log:    logrus.NewEntry(logrus.New()),
	}

}

func (t *OpensearchHandlerTestSuite) BeforeTest(suiteName, testName string) {
	httpmock.Reset()
}
