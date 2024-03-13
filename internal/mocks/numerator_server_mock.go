package mock

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/labstack/echo/v4"
)

type AbstractMockNumeratorServer struct {
	server   *echo.Echo
	recorder *httptest.ResponseRecorder
	config   *config.NumeratorConfig
}

func (m *AbstractMockNumeratorServer) Setup() {
	m.server = echo.New()
	m.recorder = httptest.NewRecorder()
	m.config = &config.NumeratorConfig{
		BaseURL: "http://localhost:8080", // Assuming the mock server runs on port 8080
		APIKey:  "test-api-key",
	}
}

func (m *AbstractMockNumeratorServer) Teardown() {
	// Clean up resources if needed
}

func (m *AbstractMockNumeratorServer) PrepareEndpoint(response []byte, pathOfJSON string, code int) string {
	var body []byte
	if response == nil && pathOfJSON != "" {
		// Read JSON data from file
		jsonFilePath := filepath.Join("testdata", "response", pathOfJSON)
		jsonFileData, err := ioutil.ReadFile(jsonFilePath)
		if err != nil {
			panic(err) // You may want to handle this error differently
		}
		body = jsonFileData
	} else if response != nil {
		body = response
	} else {
		panic("Both response and pathOfJSON cannot be empty")
	}

	// Set up route handler for mock server
	m.server.GET("/", func(c echo.Context) error {
		return c.String(code, string(body))
	})

	// Make a request to the mock server to trigger the route handler
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	m.server.ServeHTTP(m.recorder, req)

	return string(body)
}

func TestMockServer(t *testing.T) {
	mockServer := &AbstractMockNumeratorServer{}
	mockServer.Setup()
	defer mockServer.Teardown()

	responseBody := []byte(`{"message": "test response"}`)
	mockServer.PrepareEndpoint(responseBody, "test.json", http.StatusOK)

	if mockServer.recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, mockServer.recorder.Code)
	}
}
