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
	assert.Equal(suite.T(), int64(2), resp.Count())
}

func (suite *NumeratorRouterSuite) TestNumeratorRouter_FlagValueByKey() {
	flagKey := "go_featureflag_01"
	context := make(map[string]interface{})
	FlagValueByKeyRequest := request.FlagValueByKeyRequest{
		Key:     flagKey,
		Context: context,
	}

	resp := RequestSuccess[response.FeatureFlagVariationValue](suite.e, http.MethodPost, "/api/sdk/feature-flag/by-key", ToJsonString(FlagValueByKeyRequest))
	fmt.Println(resp)
	assert.Equal(suite.T(), true, resp.Value.BooleanValue)
	assert.Equal(suite.T(), enum.BOOLEAN, resp.ValueType)
}
