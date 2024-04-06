package clients

import (
	"testing"

	mock_service "github.com/c0x12c/numerator-go-sdk/internal/service/mock_service/mock_numerator_service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	mock_context "github.com/c0x12c/numerator-go-sdk/pkg/context/mock_context/mock_context_provider"
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
		response *response.SuccessResponse[response.FeatureFlagList]
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
				response: &response.SuccessResponse[response.FeatureFlagList]{
					SuccessResponse: response.FeatureFlagList{
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
				defaultValue: int64(100),
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
							LongValue: int64(100),
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
			gotValue, err := numeratorClient.FlagValueByKey(c.in.requestBody.Key, c.in.requestBody.Context)
			assert.NoError(t, err)

			// Assert that the value returned by GetValueByKeyWithDefault can be typecasted to bool
			flagExpected := expectedResponse.SuccessResponse

			switch c.in.defaultValue.(type) {
			case string:
				assert.Equal(t, flagExpected.Value.StringValue, gotValue.Value.StringValue)
			case bool:
				assert.Equal(t, flagExpected.Value.BooleanValue, gotValue.Value.BooleanValue)
			case int64:
				assert.Equal(t, flagExpected.Value.LongValue, gotValue.Value.LongValue)
			case float64:
				assert.Equal(t, flagExpected.Value.DoubleValue, gotValue.Value.DoubleValue)
			default:
				return
			}
		})
	}
}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagVariationDetail_With_No_Context_Success() {
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
			name: "fetch BOOLEAN variation value - return correct value",
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
			name: "fetch STRING variation value - return correct value",
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
			name: "fetch LONG variation value - return correct value",
			in: testCaseIn{
				defaultValue: int64(100),
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
							LongValue: int64(100),
						},
						ValueType: "LONG",
					},
				},
			},
		},
		{
			name: "fetch DOUBLE variation value - return correct value",
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

			// Expected flag
			flagExpected := expectedResponse.SuccessResponse

			switch value := c.in.defaultValue.(type) {
			case string:
				defaultString := value
				// Call the function being tested
				gotValue, err := numeratorClient.StringFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaultString, false)
				assert.NoError(t, err)
				assert.Equal(t, flagExpected.Value.StringValue, gotValue.Value)
			case bool:
				defaultBoolean := value
				// Call the function being tested
				gotValue, err := numeratorClient.BooleanFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaultBoolean, false)
				assert.NoError(t, err)
				assert.Equal(t, flagExpected.Value.BooleanValue, gotValue.Value)
			case int64:
				defaultLong := value
				// Call the function being tested
				gotValue, err := numeratorClient.LongFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaultLong, false)
				assert.NoError(t, err)
				assert.Equal(t, flagExpected.Value.LongValue, gotValue.Value)
			case float64:
				defaultDouble := value
				// Call the function being tested
				gotValue, err := numeratorClient.DoubleFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaultDouble, false)
				assert.NoError(t, err)
				assert.Equal(t, flagExpected.Value.DoubleValue, gotValue.Value)
			default:
				return
			}
		})
	}
}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagVariationDetail_With_Context_Success() {
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
			name: "fetch BOOLEAN variation value with default context - return correct value",
			in: testCaseIn{
				defaultValue: true,
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Boolean,
					Context: map[string]interface{}{"value": false},
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_01",
						Status: "ON",
						Value: response.VariationValue{
							BooleanValue: false,
						},
						ValueType: "BOOLEAN",
					},
				},
			},
		},
		{
			name: "fetch STRING variation value with default context - return correct value",
			in: testCaseIn{
				defaultValue: "one",
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_String,
					Context: map[string]interface{}{"value": "two"},
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_02",
						Status: "ON",
						Value: response.VariationValue{
							StringValue: "two",
						},
						ValueType: "STRING",
					},
				},
			},
		},
		{
			name: "fetch LONG variation value with default context - return correct value",
			in: testCaseIn{
				defaultValue: int64(100),
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Long,
					Context: map[string]interface{}{"value": int64(200)},
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_03",
						Status: "ON",
						Value: response.VariationValue{
							LongValue: int64(200),
						},
						ValueType: "LONG",
					},
				},
			},
		},
		{
			name: "fetch DOUBLE variation value with default context - return correct value",
			in: testCaseIn{
				defaultValue: 1.5,
				requestBody: request.FlagValueByKeyRequest{
					Key:     constant.FlagKey_Double,
					Context: map[string]interface{}{"value": 2.5},
				},
			},
			expected: testCaseOut{
				response: &response.SuccessResponse[response.FeatureFlagVariationValue]{
					SuccessResponse: response.FeatureFlagVariationValue{
						Key:    "go_featureflag_04",
						Status: "ON",
						Value: response.VariationValue{
							DoubleValue: 2.5,
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

			// Create mock instances
			mockNumeratorService := mock_service.NewMockNumeratorService(ctrl)
			mockContextProvider := mock_context.NewMockContextProvider(ctrl)

			// Define expected response from the mock service
			expectedResponse := c.expected.response

			// Set up expectations on the mock service
			mockNumeratorService.EXPECT().FlagValueByKey(c.in.requestBody).Return(expectedResponse, nil)
			mockContextProvider.EXPECT().Context().Return(c.in.requestBody.Context)

			// Create the NumeratorClient with the mock service
			numeratorClient := &DefaultNumeratorClient{
				service:         mockNumeratorService,
				contextProvider: mockContextProvider,
			}

			// Expected flag
			flagExpected := expectedResponse.SuccessResponse

			switch value := c.in.defaultValue.(type) {
			case string:
				defaultString := value
				// Call the function being tested
				gotValue, err := numeratorClient.StringFlagVariationDetail(c.in.requestBody.Key, nil, defaultString, true)
				assert.NoError(t, err)
				assert.Empty(t, gotValue.Reason)
				assert.Equal(t, flagExpected.Value.StringValue, gotValue.Value)
			case bool:
				defaultBoolean := value
				// Call the function being tested
				gotValue, err := numeratorClient.BooleanFlagVariationDetail(c.in.requestBody.Key, nil, defaultBoolean, true)
				assert.NoError(t, err)
				assert.Empty(t, gotValue.Reason)
				assert.Equal(t, flagExpected.Value.BooleanValue, gotValue.Value)
			case int64:
				defaultLong := value
				// Call the function being tested
				gotValue, err := numeratorClient.LongFlagVariationDetail(c.in.requestBody.Key, nil, defaultLong, true)
				assert.NoError(t, err)
				assert.Empty(t, gotValue.Reason)
				assert.Equal(t, flagExpected.Value.LongValue, gotValue.Value)
			case float64:
				defaultDouble := value
				// Call the function being tested
				gotValue, err := numeratorClient.DoubleFlagVariationDetail(c.in.requestBody.Key, nil, defaultDouble, true)
				assert.NoError(t, err)
				assert.Empty(t, gotValue.Reason)
				assert.Equal(t, flagExpected.Value.DoubleValue, gotValue.Value)
			default:
				return
			}
		})
	}
}

func (s *testNumeratorClientSuite) TestFeatureFlag_FlagVariationDetail_Failure() {
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
			name: "fetch BOOLEAN variation value - return default value",
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
						ValueType: "STRING",
					},
				},
			},
		},
		{
			name: "fetch STRING variation value - return default value",
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
						ValueType: "BOOLEAN",
					},
				},
			},
		},
		{
			name: "fetch LONG variation value - return default value",
			in: testCaseIn{
				defaultValue: int64(100),
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
							LongValue: int64(100),
						},
						ValueType: "DOUBLE",
					},
				},
			},
		},
		{
			name: "fetch DOUBLE variation value - return default value",
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
						ValueType: "LONG",
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

			expectedReason := map[string]interface{}{
				"kind":      "Error",
				"errorKind": "type mismatch",
			}

			switch defaulVal := c.in.defaultValue.(type) {
			case string:
				// Call the function being tested
				gotValue, err := numeratorClient.StringFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaulVal, false)
				assert.NoError(t, err)
				assert.Equal(t, defaulVal, gotValue.Value)
				assert.Equal(t, expectedReason, gotValue.Reason)
			case bool:
				// Call the function being tested
				gotValue, err := numeratorClient.BooleanFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaulVal, false)
				assert.NoError(t, err)
				assert.Equal(t, defaulVal, gotValue.Value)
				assert.Equal(t, expectedReason, gotValue.Reason)
			case int64:
				// Call the function being tested
				gotValue, err := numeratorClient.LongFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaulVal, false)
				assert.NoError(t, err)
				assert.Equal(t, defaulVal, gotValue.Value)
				assert.Equal(t, expectedReason, gotValue.Reason)
			case float64:
				// Call the function being tested
				gotValue, err := numeratorClient.DoubleFlagVariationDetail(c.in.requestBody.Key, c.in.requestBody.Context, defaulVal, false)
				assert.NoError(t, err)
				assert.Equal(t, defaulVal, gotValue.Value)
				assert.Equal(t, expectedReason, gotValue.Reason)
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
			_, gotError := numeratorClient.FlagValueByKey(c.in.requestBody.Key, c.in.requestBody.Context)
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
