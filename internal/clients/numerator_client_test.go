package clients

import (
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/config"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type testNumeratorClientSuite struct {
	suite.Suite
}

func TestNumeratorClientSuite(t *testing.T) {
	suite.Run(t, &testNumeratorClientSuite{})
}

func (s *testNumeratorClientSuite) TestFeatureFlags() {
	type testCaseIn struct {
		page int
		size int
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
				page: 0,
				size: 1,
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

	//setup dependencies for test
	numeratorConfig := config.NewNumeratorConfig(constant.ApiKeyTest)
	numeratorClient := NewDefaultNumeratorClient(numeratorConfig)

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			flags, err := numeratorClient.FeatureFlags(c.in.page, c.in.size)

			flagExpected := c.expected.featureFlags[0]
			flagActual := flags[0]

			assert.NoError(t, err)
			assert.Equal(t, c.in.size, len(flags))
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
