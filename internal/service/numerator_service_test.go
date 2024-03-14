package service

import (
	"testing"

	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type testNumeratorServiceSuite struct {
	suite.Suite
}

func TestNumeratorServiceSuite(t *testing.T) {
	suite.Run(t, &testNumeratorServiceSuite{})
}

func (s *testNumeratorServiceSuite) TestFeatureFlags_FlagListSuccess() {
	type testCaseIn struct {
		requestBody request.FlagListRequest
	}

	type testCaseOut struct {
		featureFlags []response.FeatureFlag
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "fetch data from numerator successfully",
			in: testCaseIn{
				requestBody: request.FlagListRequest{
					Page: 0,
					Size: 1,
				},
			},
			expected: testCaseOut{
				featureFlags: []response.FeatureFlag{
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
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// // Create a mock Service
			// mockNumeratorService := mock_service.NewMockNumeratorService(ctrl)

			// // Define expectations for the FeatureFlags method
			// mockNumeratorService.EXPECT().FlagList(c.in.requestBody).Return(c.expected.featureFlags, nil)

			// // Call the function being tested
			// expectedFlags, err := mockNumeratorService.FeatureFlags(c.in.page, c.in.size)
			// assert.NoError(t, err)

			// gotFlags, err := numeratorService.FeatureFlags(c.in.page, c.in.size)
			// assert.NoError(t, err)

			// flagExpected := expectedFlags[0]
			// flagActual := gotFlags[0]

			// assert.Equal(t, c.in.size, len(gotFlags))
			// assert.Equal(t, flagExpected.Name, flagActual.Name)
			// assert.Equal(t, flagExpected.Status, flagActual.Status)
			// assert.Equal(t, flagExpected.Key, flagActual.Key)
			// assert.Equal(t, flagExpected.Description, flagActual.Description)
			// assert.Equal(t, flagExpected.DefaultOnVariation, flagActual.DefaultOnVariation)
			// assert.Equal(t, flagExpected.DefaultOffVariation, flagActual.DefaultOffVariation)
			// assert.Equal(t, flagExpected.ValueType, flagActual.ValueType)
		})
	}

}
