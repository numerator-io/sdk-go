package clients

import (
	"testing"

	mock_service "github.com/c0x12c/numerator-go-sdk/internal/service/mock_service/mock_numerator_service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	"github.com/c0x12c/numerator-go-sdk/pkg/exception"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testNumeratorClientSuite struct {
	suite.Suite
}

func TestNumeratorClientSuite(t *testing.T) {
	suite.Run(t, &testNumeratorClientSuite{})
}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagListSuccess() {
	type testCaseIn struct {
		requestBody request.FlagListRequest
	}

	type testCaseOut struct {
		response *response.SuccessResponse[response.FeatureFlagListResponse]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "fetch flag list - success",
			in: testCaseIn{
				requestBody: request.FlagListRequest{
					Page: 0,
					Size: 1,
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagListResponse]{
					SuccessResponse: response.FeatureFlagListResponse{
						CountVal: 1,
						DataVal: []response.FeatureFlag{
							{
								Name:        "Go Numerator Test",
								Key:         "go_featureflag_01",
								Status:      "ON",
								Description: "test go_featureflag_01",
								DefaultOnVariation: response.FlagVariation{
									Name: "default_on",
									Value: response.VariationValue{
										BooleanValue: true,
									},
								},
								DefaultOffVariation: response.FlagVariation{
									Name:  "default_off",
									Value: response.VariationValue{},
								},
								ValueType: "BOOLEAN",
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorService.EXPECT().FlagList(c.in.requestBody).Return(expectedResponse, nil)

			numeratorClient := &DefaultNumeratorClient{
				service: mockNumeratorService,
			}

			// Call the function being tested
			gotFlags, err := numeratorClient.FeatureFlags(c.in.requestBody.Page, c.in.requestBody.Size)
			assert.NoError(t, err)

			flagDataExpected := expectedResponse.SuccessResponse.Data()
			flagExpected := flagDataExpected[0]

			flagActual := gotFlags[0]

			assert.Equal(t, c.in.requestBody.Size, len(flagDataExpected))
			assert.Equal(t, flagExpected.Name, flagActual.Name)
			assert.Equal(t, flagExpected.Status, flagActual.Status)
			assert.Equal(t, flagExpected.Key, flagActual.Key)
			assert.Equal(t, flagExpected.Description, flagActual.Description)
			assert.Equal(t, flagExpected.DefaultOnVariation, flagActual.DefaultOnVariation)
			assert.Equal(t, flagExpected.DefaultOffVariation, flagActual.DefaultOffVariation)
			assert.Equal(t, flagExpected.ValueType, flagActual.ValueType)
		})
	}

}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagValueByKey_Success() {
	type testCaseIn struct {
		defaultValue interface{}
		requestBody  request.FlagValueByKeyRequest
	}

	type testCaseOut struct {
		response *response.SuccessResponse[response.FeatureFlagVariationValue]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "fetch flag value by BOOLEAN key - return correct value",
			in: testCaseIn{
				defaultValue: true,
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Boolean,
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_01",
						Status: "ON",
						Value: response.VariationValue{
							BooleanValue: true,
						},
						ValueType: "BOOLEAN",
					},
				},
			},
		},
		{
			name: "fetch flag value by STRING key - return correct value",
			in: testCaseIn{
				defaultValue: "default_on",
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_String,
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_02",
						Status: "ON",
						Value: response.VariationValue{
							StringValue: "default_on",
						},
						ValueType: "STRING",
					},
				},
			},
		},
		{
			name: "fetch flag value by LONG key - return correct value",
			in: testCaseIn{
				defaultValue: 100,
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Long,
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_03",
						Status: "ON",
						Value: response.VariationValue{
							LongValue: 100,
						},
						ValueType: "LONG",
					},
				},
			},
		},
		{
			name: "fetch flag value by DOUBLE key - return correct value",
			in: testCaseIn{
				defaultValue: 1.5,
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Double,
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_04",
						Status: "ON",
						Value: response.VariationValue{
							DoubleValue: 1.5,
						},
						ValueType: "DOUBLE",
					},
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock NumeratorService
			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl)

			// Define expected response from the mock service
			expectedResponse := c.expected.response

			// Set up expectations on the mock service
			mockNumeratorService.EXPECT().FlagValueByKey(c.in.requestBody).Return(expectedResponse, nil)

			// Create the NumeratorClient with the mock service
			numeratorClient := &DefaultNumeratorClient{
				service: mockNumeratorService,
			}

			// Call the function being tested
			gotValue, err := numeratorClient.GetValueByKeyWithDefault(c.in.requestBody.Key, c.in.requestBody.Context, c.in.defaultValue)
			assert.NoError(t, err)

			// Assert that the value returned by GetValueByKeyWithDefault can be typecasted to bool
			flagExpected := expectedResponse.SuccessResponse
			var extractVal interface{}
			var ok bool

			switch c.in.defaultValue.(type) {
			case string:
				extractVal, ok = gotValue.(string)
				assert.True(t, ok, "returned value is not of type STRING")
				assert.Equal(t, flagExpected.Value.StringValue, extractVal)

			case bool:
				extractVal, ok = gotValue.(bool)
				assert.True(t, ok, "returned value is not of type BOOLEAN")
				assert.Equal(t, flagExpected.Value.BooleanValue, extractVal)
			case int:
				extractVal, ok = gotValue.(int64)
				assert.True(t, ok, "returned value is not of type LONG")
				assert.Equal(t, flagExpected.Value.LongValue, extractVal)
			case float64:
				extractVal, ok = gotValue.(float64)
				assert.True(t, ok, "returned value is not of type DOUBLE")
				assert.Equal(t, flagExpected.Value.DoubleValue, extractVal)
			default:
				return
			}
		})
	}
}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagValueByKey_Failure() {
	type testCaseIn struct {
		defaultValue interface{}
		requestBody  request.FlagValueByKeyRequest
	}

	type testCaseOut struct {
		response *response.FailureResponse
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get variation value not found - failure",
			in: testCaseIn{
				defaultValue: true,
				requestBody: request.FlagValueByKeyRequest{
					Key:     "key_not_found",
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.FailureResponse{
					Error: response.NumeratorError{
						Message:    "FLAG_NOT_FOUND",
						HttpStatus: 404,
					},
				},
			},
		},
		{
			name: "get variation value mismatch type - failure",
			in: testCaseIn{
				defaultValue: true,
				requestBody: request.FlagValueByKeyRequest{
					Key:     "key_not_found",
					Context: nil,
				},
			},
			expected: testCaseOut{
				response: &response.FailureResponse{
					Error: response.NumeratorError{
						Message:    "type mismatch",
						HttpStatus: 401,
					},
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock NumeratorService
			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl)

			// Define expected response from the mock service
			expectedResponse := c.expected.response

			// Set up expectations on the mock service
			mockNumeratorService.EXPECT().FlagValueByKey(c.in.requestBody).Return(expectedResponse, nil)

			// Create the NumeratorClient with the mock service
			numeratorClient := &DefaultNumeratorClient{
				service: mockNumeratorService,
			}

			// Call the function being tested
			_, gotError := numeratorClient.GetValueByKeyWithDefault(c.in.requestBody.Key, c.in.requestBody.Context, c.in.defaultValue)
			extractError, ok := gotError.(*exception.NumeratorException)
			assert.True(t, ok, "returned value is not of type NumeratorException")

			errorExpected := expectedResponse.Error

			assert.Equal(t, errorExpected.Message, extractError.Message)
			assert.Equal(t, errorExpected.HttpStatus, extractError.Status)
		})
	}

}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagDetailByKey_Success() {
	type testCaseIn struct {
		flagKey string
	}

	type testCaseOut struct {
		response *response.SuccessResponse[response.FeatureFlag]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "fetch detail flag by BOOLEAN key - success",
			in: testCaseIn{
				flagKey: constant.FlagKey_Boolean,
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlag]{
					SuccessResponse: response.FeatureFlag{
						Name:        "Go Numerator Test",
						Key:         "go_featureflag_01",
						Status:      "ON",
						Description: "test go_featureflag_01",
						DefaultOnVariation: response.FlagVariation{
							Name: "default_on",
							Value: response.VariationValue{
								BooleanValue: true,
							},
						},
						DefaultOffVariation: response.FlagVariation{
							Name:  "default_off",
							Value: response.VariationValue{},
						},
						ValueType: "BOOLEAN",
					},
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorService.EXPECT().FlagDetailByKey(c.in.flagKey).Return(expectedResponse, nil)

			numeratorClient := &DefaultNumeratorClient{
				service: mockNumeratorService,
			}

			// Call the function being tested
			flagActual, err := numeratorClient.FeatureFlagDetails(c.in.flagKey)
			assert.NoError(t, err)

			flagExpected := expectedResponse.SuccessResponse

			assert.Equal(t, flagExpected.Name, flagActual.Name)
			assert.Equal(t, flagExpected.Status, flagActual.Status)
			assert.Equal(t, flagExpected.Key, flagActual.Key)
			assert.Equal(t, flagExpected.Description, flagActual.Description)
			assert.Equal(t, flagExpected.DefaultOnVariation, flagActual.DefaultOnVariation)
			assert.Equal(t, flagExpected.DefaultOffVariation, flagActual.DefaultOffVariation)
			assert.Equal(t, flagExpected.ValueType, flagActual.ValueType)
		})
	}

}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagDetailByKey_Failure() {
	type testCaseIn struct {
		flagKey string
	}

	type testCaseOut struct {
		response *response.FailureResponse
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get variation value not found - failure",
			in: testCaseIn{
				flagKey: "key_not_found",
			},
			expected: testCaseOut{
				response: &response.FailureResponse{
					Error: response.NumeratorError{
						Message:    "FLAG_NOT_FOUND",
						HttpStatus: 404,
					},
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorService.EXPECT().FlagDetailByKey(c.in.flagKey).Return(expectedResponse, nil)

			numeratorClient := &DefaultNumeratorClient{
				service: mockNumeratorService,
			}

			// Call the function being tested
			_, gotError := numeratorClient.FeatureFlagDetails(c.in.flagKey)
			extractError, ok := gotError.(*exception.NumeratorException)
			assert.True(t, ok, "returned value is not of type NumeratorException")

			errorExpected := expectedResponse.Error

			assert.Equal(t, errorExpected.Message, extractError.Message)
			assert.Equal(t, errorExpected.HttpStatus, extractError.Status)
		})
	}

}