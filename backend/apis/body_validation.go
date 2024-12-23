package apis

import (
	"errors"
	"vending_machine/dtos"

	"github.com/go-playground/validator/v10"
)

var val = validator.New()
 
func Validation(i any) []*dtos.ErrorValidation {
	var errs []*dtos.ErrorValidation
 
	if err := val.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, e := range ve {
				var element dtos.ErrorValidation
				element.Tag = e.Tag()
				element.Field = e.Field()
				element.Message = "Error:Field validation for '" + e.Field() + "' failed on the '" + e.Tag() + "' tag"
				errs = append(errs, &element)
			}
		}
	}

	return errs
}