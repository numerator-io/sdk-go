package route

import (
	"github.com/c0x12c/numerator-go-sdk/internal/service"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/request"
	"github.com/c0x12c/numerator-go-sdk/pkg/api/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type numeratorRouter struct {
	numeratorService service.NumeratorService
	validate         *validator.Validate
}

func (r *numeratorRouter) Configure(e *echo.Echo) {
	numGroup := e.Group("/api/sdk/feature-flag")

	numGroup.POST("/by-key", r.FlagValueByKey)
	numGroup.POST("/detail-by-key", r.FlagDetailByKey)
	numGroup.POST("/listing", r.FlagList)
}

func NewNumeratorRouter(numeratorService service.NumeratorService) *numeratorRouter {
	validate := validator.New()
	return &numeratorRouter{
		numeratorService: numeratorService,
		validate:         validate,
	}
}

// FlagValueByKey godoc
//
// @Summary			API for retrieving a feature flag value by key
// @Description		API for retrieving a feature flag value by key along with a context
// @Produce			json
// @Success			200	{object}	response.SuccessResponse
// @Failure			401	{object}	response.ErrorResponse
// @Failure			500	{object}	response.ErrorResponse
// @Router				/api/sdk/feature-flag/by-key [POST]
// @Tags				FeatureFlag
// @Security			BearerAuth
// @Param				body request.FlagValueByKeyRequest true "body"
func (r *numeratorRouter) FlagValueByKey(c echo.Context) (err error) {
	return HandleRequestWithBody[request.FlagValueByKeyRequest, response.ApiResponse](
		c,
		r.validate,
		func(req request.FlagValueByKeyRequest) (response.ApiResponse, error) {
			return r.numeratorService.FlagValueByKey(req)
		},
	)
}

// FlagDetailByKey godoc
//
// @Summary			API for retrieving detailed information about a feature flag by key
// @Description		API for retrieving detailed information about a feature flag by key
// @Produce			json
// @Success			200	{object}	response.SuccessResponse
// @Failure			401	{object}	response.ErrorResponse
// @Failure			500	{object}	response.ErrorResponse
// @Router				/api/sdk/feature-flag/detail-by-key [POST]
// @Tags				FeatureFlag
// @Security			BearerAuth
// @Param				key query string true "flag key"
func (r *numeratorRouter) FlagDetailByKey(c echo.Context) (err error) {
	return HandleRequestWithBody[string, response.ApiResponse](
		c,
		r.validate,
		func(flagKey string) (response.ApiResponse, error) {
			return r.numeratorService.FlagDetailByKey(flagKey)
		},
	)
}

// FlagList godoc
//
// @Summary			API for retrieving a list of feature flags
// @Description		API for retrieving a list of feature flags with pagination support
// @Produce			json
// @Success			200	{object}	response.SuccessResponse
// @Failure			401	{object}	response.ErrorResponse
// @Failure			500	{object}	response.ErrorResponse
// @Router				/api/sdk/feature-flag/listing [POST]
// @Tags				FeatureFlag
// @Security			BearerAuth
// @Param				body request.FlagListRequest true "body"
func (r *numeratorRouter) FlagList(c echo.Context) (err error) {
	return HandleRequestWithBody[request.FlagListRequest, response.ApiResponse](
		c,
		r.validate,
		func(req request.FlagListRequest) (response.ApiResponse, error) {
			return r.numeratorService.FlagList(req)
		},
	)
}
