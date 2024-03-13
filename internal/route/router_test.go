package route_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/c0x12c/numerator-go-sdk/internal/route"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
	e        *echo.Echo
	ctx      context.Context
	startAt  time.Time
	services []any
}

func TestRunRouterSuites(t *testing.T) {
	suite.Run(t, new(NumeratorRouterSuite))
}

// SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *RouterSuite) SetupSuite() {
	suite.startAt = time.Now()
	suite.ctx = context.TODO()
	suite.e = echo.New()

	mockInstance := []any{}
	routers, services, _ := route.Routers(suite.ctx, mockInstance)
	suite.services = services
	for _, router := range routers {
		router.Configure(suite.e)
	}
}

func (suite *RouterSuite) TearDownSuite() {

}

// run after each test
func (s *RouterSuite) TearDownTest() {
	s.TearDownSuite()
}

func Request(e *echo.Echo, method string, target string, bodyStr string) []byte {
	var body io.Reader
	if bodyStr != "" {
		body = bytes.NewBufferString(bodyStr)
	} else {
		body = nil
	}
	path := constant.BaseURL + target
	connectRequest := httptest.NewRequest(method, path, body)
	connectRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	connectRequest.Header.Add(echo.HeaderAccept, echo.MIMEApplicationJSON)
	apiKey := os.Getenv("API_KEY")
	if apiKey != "" {
		connectRequest.Header.Add(constant.XNumAPIKeyHeader, apiKey)
	}
	recorder := httptest.NewRecorder()
	e.ServeHTTP(recorder, connectRequest)
	return recorder.Body.Bytes()
}

func RequestSuccess[T interface{}](e *echo.Echo, method string, target string, bodyStr string) T {
	var res response.SuccessResponse[T]
	err := json.Unmarshal(Request(e, method, target, bodyStr), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res.SuccessResponse
}

func ToJsonString(data any) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
