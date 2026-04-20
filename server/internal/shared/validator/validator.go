package validator

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)

			for _, e := range errs {
				out[e.Field()] = e.Tag()
			}

			outJSON, _ := json.Marshal(out)

			return &echo.HTTPError{
				Code:    400,
				Message: string(outJSON),
			}
		}

		return echo.ErrBadRequest
	}
	return nil
}
