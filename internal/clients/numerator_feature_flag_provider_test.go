package clients

import (
	"testing"

	mock_clients "github.com/c0x12c/numerator-go-sdk/internal/clients/mock_client/mock_numerator_client"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testNumeratorFeatureFlagProviderSuite struct {
	suite.Suite
}

func TestNumeratorFeatureFlagProviderSuite(t *testing.T) {
	suite.Run(t, &testNumeratorFeatureFlagProviderSuite{})
}

func (s *testNumeratorFeatureFlagProviderSuite) TestGetBooleanFeatureFlag_Success() {
	type testCaseIn struct {
		flagKey      string
		defaultValue bool
	}

	type testCaseOut struct {
		response *response.FlagEvaluationDetail[bool]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get feature flag - success",
			in: testCaseIn{
				flagKey:      constant.FlagKey_Boolean,
				defaultValue: true,
			},
			expected: testCaseOut{
				response: &response.FlagEvaluationDetail[bool]{
					Key:    "go_feature_flag_01",
					Value:  false,
					Reason: nil,
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorClient := mock_clients.NewMockNumeratorClient(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorClient.EXPECT().BooleanFlagVariationDetail(c.in.flagKey, nil, c.in.defaultValue, false).Return(expectedResponse, nil)

			numeratorFFProvider := &NumeratorFeatureFlagProvider{
				client: mockNumeratorClient,
			}

			// Call the function being tested
			gotValue := numeratorFFProvider.GetBooleanFeatureFlag(c.in.flagKey, c.in.defaultValue, nil, false)
			expectedValue := expectedResponse.Value

			assert.Equal(t, expectedValue, gotValue)
		})
	}

}

func (s *testNumeratorFeatureFlagProviderSuite) TestGetStringFeatureFlag_Success() {
	type testCaseIn struct {
		flagKey      string
		defaultValue string
	}

	type testCaseOut struct {
		response *response.FlagEvaluationDetail[string]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get feature flag - success",
			in: testCaseIn{
				flagKey:      constant.FlagKey_String,
				defaultValue: "one",
			},
			expected: testCaseOut{
				response: &response.FlagEvaluationDetail[string]{
					Key:    "go_feature_flag_02",
					Value:  "two",
					Reason: nil,
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorClient := mock_clients.NewMockNumeratorClient(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorClient.EXPECT().StringFlagVariationDetail(c.in.flagKey, nil, c.in.defaultValue, false).Return(expectedResponse, nil)

			numeratorFFProvider := &NumeratorFeatureFlagProvider{
				client: mockNumeratorClient,
			}

			// Call the function being tested
			gotValue := numeratorFFProvider.GetStringFeatureFlag(c.in.flagKey, c.in.defaultValue, nil, false)
			expectedValue := expectedResponse.Value

			assert.Equal(t, expectedValue, gotValue)
		})
	}

}

func (s *testNumeratorFeatureFlagProviderSuite) TestGetLongFeatureFlag_Success() {
	type testCaseIn struct {
		flagKey      string
		defaultValue int64
	}

	type testCaseOut struct {
		response *response.FlagEvaluationDetail[int64]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get feature flag - success",
			in: testCaseIn{
				flagKey:      constant.FlagKey_Long,
				defaultValue: int64(100),
			},
			expected: testCaseOut{
				response: &response.FlagEvaluationDetail[int64]{
					Key:    "go_feature_flag_03",
					Value:  int64(200),
					Reason: nil,
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorClient := mock_clients.NewMockNumeratorClient(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorClient.EXPECT().LongFlagVariationDetail(c.in.flagKey, nil, c.in.defaultValue, false).Return(expectedResponse, nil)

			numeratorFFProvider := &NumeratorFeatureFlagProvider{
				client: mockNumeratorClient,
			}

			// Call the function being tested
			gotValue := numeratorFFProvider.GetLongFeatureFlag(c.in.flagKey, c.in.defaultValue, nil, false)
			expectedValue := expectedResponse.Value

			assert.Equal(t, expectedValue, gotValue)
		})
	}

}

func (s *testNumeratorFeatureFlagProviderSuite) TestGetDoubleFeatureFlag_Success() {
	type testCaseIn struct {
		flagKey      string
		defaultValue float64
	}

	type testCaseOut struct {
		response *response.FlagEvaluationDetail[float64]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get feature flag - success",
			in: testCaseIn{
				flagKey:      constant.FlagKey_Double,
				defaultValue: 1.5,
			},
			expected: testCaseOut{
				response: &response.FlagEvaluationDetail[float64]{
					Key:    "go_feature_flag_04",
					Value:  2.5,
					Reason: nil,
				},
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockNumeratorClient := mock_clients.NewMockNumeratorClient(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorClient.EXPECT().DoubleFlagVariationDetail(c.in.flagKey, nil, c.in.defaultValue, false).Return(expectedResponse, nil)

			numeratorFFProvider := &NumeratorFeatureFlagProvider{
				client: mockNumeratorClient,
			}

			// Call the function being tested
			gotValue := numeratorFFProvider.GetDoubleFeatureFlag(c.in.flagKey, c.in.defaultValue, nil, false)
			expectedValue := expectedResponse.Value

			assert.Equal(t, expectedValue, gotValue)
		})
	}

}

func (s *testNumeratorFeatureFlagProviderSuite) TestGetBooleanFeatureFlag_Failure() {
	type testCaseIn struct {
		flagKey      string
		defaultValue bool
	}

	type testCaseOut struct {
		response *response.FlagEvaluationDetail[bool]
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "get feature flag - failure",
			in: testCaseIn{
				flagKey:      constant.FlagKey_Boolean,
				defaultValue: true,
			},
			expected: testCaseOut{
				response: &response.FlagEvaluationDetail[bool]{
					Key:   "go_feature_flag_01",
					Value: true,
					Reason: map[string]interface{}{
						"kind":      "Error",
						"errorKind": "type mismatch",
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

			mockNumeratorClient := mock_clients.NewMockNumeratorClient(ctrl) // Create a mock client

			expectedResponse := c.expected.response
			// Define expectations for the FeatureFlags method
			mockNumeratorClient.EXPECT().BooleanFlagVariationDetail(c.in.flagKey, nil, c.in.defaultValue, false).Return(expectedResponse, nil)

			numeratorFFProvider := &NumeratorFeatureFlagProvider{
				client: mockNumeratorClient,
			}

			// Call the function being tested
			gotValue := numeratorFFProvider.GetBooleanFeatureFlag(c.in.flagKey, c.in.defaultValue, nil, false)
			expectedValue := expectedResponse.Value

			assert.Equal(t, expectedValue, gotValue)
		})
	}

}
