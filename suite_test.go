package opensearchhandler

import (
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

	client, err := opensearch.NewClient(opensearch.SetURL(baseURL), opensearch.SetTransport(httpmock.DefaultTransport), opensearch.SetHealthcheck(false), opensearch.SetSniff(false))
	if err != nil {
		panic(err)
	}

	httpmock.Activate()

	t.opensearchHandler = &OpensearchHandlerImpl{
		client: client,
		log:    logrus.NewEntry(logrus.New()),
	}

}

func (t *OpensearchHandlerTestSuite) BeforeTest(suiteName, testName string) {
	httpmock.Reset()
}
