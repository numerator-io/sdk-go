package route_test

import (
	"fmt"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"
	"github.com/c0x12c/numerator-go-sdk/pkg/constant"
	"github.com/c0x12c/numerator-go-sdk/pkg/enum"
	"github.com/stretchr/testify/assert"
)

type NumeratorRouterSuite struct {
	RouterSuite
}

func (suite *NumeratorRouterSuite) TestNumeratorRouter_FlagList() {
	flagListRequest := request.FlagListRequest{
		Page: constant.Page,
		Size: constant.Size,
	}

	resp := RequestSuccess[response.FeatureFlagListResponse](suite.e, http.MethodPost, "/api/sdk/feature-flag/listing", ToJsonString(flagListRequest))
	assert.Equal(suite.T(), int64(4), resp.Count())
}

func (suite *NumeratorRouterSuite) TestNumeratorRouter_FlagValueByKey_TestBooleanValue() {
	flagKey := constant.FlagKey_Boolean
	context := make(map[string]interface{})
	FlagValueByKeyRequest := request.FlagValueByKeyRequest{
		Key:     flagKey,
		Context: context,
	}

	resp := RequestSuccess[response.FeatureFlagVariationValue](suite.e, http.MethodPost, "/api/sdk/feature-flag/by-key", ToJsonString(FlagValueByKeyRequest))
	assert.Equal(suite.T(), true, resp.Value.BooleanValue)
	assert.Equal(suite.T(), enum.BOOLEAN, resp.ValueType)
}

func (suite *NumeratorRouterSuite) TestNumeratorRouter_FlagValueByKey_TestStringValue() {
	flagKey := constant.FlagKey_String
	context := make(map[string]interface{})
	FlagValueByKeyRequest := request.FlagValueByKeyRequest{
		Key:     flagKey,
		Context: context,
	}

	resp := RequestSuccess[response.FeatureFlagVariationValue](suite.e, http.MethodPost, "/api/sdk/feature-flag/by-key", ToJsonString(FlagValueByKeyRequest))
	assert.Equal(suite.T(), "off", resp.Value.StringValue)
	assert.Equal(suite.T(), enum.STRING, resp.ValueType)
}

func (suite *NumeratorRouterSuite) TestNumeratorRouter_FlagDetailByKey() {
	flagKey := constant.FlagKey_Boolean
	target := fmt.Sprintf("/api/sdk/feature-flag/detail-by-key?key=%s", flagKey)
	resp := RequestSuccess[response.FeatureFlag](suite.e, http.MethodPost, target, "")
	fmt.Println(target)
	assert.Equal(suite.T(), true, resp.DefaultOnVariation.Value.BooleanValue)
	assert.Equal(suite.T(), false, resp.DefaultOffVariation.Value.BooleanValue)
	assert.Equal(suite.T(), enum.BOOLEAN, resp.ValueType)
}
