package route

import (
	"errors"
	"net/http"

	"github.com/c0x12c/numerator-go-sdk/pkg/exception"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func HandleRequestWithBody[R, T interface{}](c echo.Context, validate *validator.Validate, execute func(req R) (out T, err error)) (err error) {
	requestBody := new(R)
	if _, err = BindAndValidate(c, *validate, requestBody); err != nil {
		return err
	}
	res, err := execute(*requestBody)
	if err != nil {
		return exception.NewNumeratorException(err.Error(), 0)
	}
	return c.JSON(http.StatusOK, res)
}

func BindAndValidate[T interface{}](c echo.Context, validate validator.Validate, requestData T) (out T, err error) {
	if err = c.Bind(requestData); err != nil {
		return out, &exception.NumeratorException{
			Status:  0,
			Message: err.Error(),
		}
	}

	//use the validator library to validate required fields
	if err = Validate(validate, requestData); err != nil {
		return out, err
	}

	return requestData, err
}

func Validate[T interface{}](validate validator.Validate, requestData T) error {
	if err := validate.Struct(requestData); err != nil {
		// Check if the error is a validation error
		var validationErrs validator.ValidationErrors
		if ok := errors.As(err, &validationErrs); ok {
			return &exception.NumeratorException{
				Status:  0,
				Message: "Validation failed",
			}
		}
	}
	return nil
}
