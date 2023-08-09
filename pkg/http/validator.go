package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

var validate = validator.New()

func ValidateStruct[k interface{}](payload k) []*ValidationDetails {
	errors := make([]*ValidationDetails, 0)

	if err := validate.Struct(payload); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			vl := ValidationDetails{
				Field:       LowerFirstChar(strcase.ToSnake(err.Field())),
				Value:       fmt.Sprintf("%v", err.Value()),
				Constraint:  err.ActualTag(),
				Description: FormatMessage(err.ActualTag()),
			}

			errors = append(errors, &vl)
		}
	}

	return errors
}
