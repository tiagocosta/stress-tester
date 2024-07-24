package stress

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	testServer *httptest.Server
}

func (suite *TestSuite) SetupSuite() {
	suite.testServer = httptest.NewServer(http.HandlerFunc(randomStatusCode))
}

func (suite *TestSuite) TearDownSuite() {
	suite.testServer.Close()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var statusCodes = [35]int{
	200, 200, 200, 200, 200, 200, 200, 200, 200, 200,
	300, 301, 302, 304, 307, 400, 401, 403, 404, 404,
	404, 404, 404, 404, 500, 500, 500, 500, 500, 501,
	503, 503, 503, 503, 503}

var index = 0

func randomStatusCode(w http.ResponseWriter, r *http.Request) {
	fooStatusCode := statusCodes[index]
	index += 1
	if index >= len(statusCodes) {
		index = 0
	}
	w.WriteHeader(fooStatusCode)
	io.WriteString(w, "HTTP Status Code")
}

func (suite *TestSuite) TestTesterStress() {
	mapStatusCode := map[int]int{
		http.StatusOK:                  10,
		http.StatusMultipleChoices:     1,
		http.StatusMovedPermanently:    1,
		http.StatusFound:               1,
		http.StatusNotModified:         1,
		http.StatusTemporaryRedirect:   1,
		http.StatusBadRequest:          1,
		http.StatusUnauthorized:        1,
		http.StatusForbidden:           1,
		http.StatusNotFound:            6,
		http.StatusInternalServerError: 5,
		http.StatusNotImplemented:      1,
		http.StatusServiceUnavailable:  5,
	}
	tester := Tester{
		URL:         fmt.Sprintf("%s/", suite.testServer.URL),
		Requests:    35,
		Concurrency: 2,
	}
	tester.Stress()
	assert.Equal(suite.T(), tester.Requests, tester.TotalRequests)
	assert.Equal(suite.T(), mapStatusCode, tester.MapStatusCode)
}
