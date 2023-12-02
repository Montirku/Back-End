package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func Validation(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" && e.Field() == "Email" {
					message = fmt.Sprintf("%s tidak boleh kosong", e.Field())
				} else if e.Tag() == "required" {
					message = fmt.Sprintf("%s tidak boleh kosong", e.Field())
				} else if e.Tag() == "email" {
					message = "email tidak valid"
				} else if e.Tag() == "max" && e.Field() == "Phone" {
					message = "nomor telepon tidak boleh lebih dari 13 digit"
				} else if e.Field() == "Phone" || e.Tag() == "min" || e.Tag() == "max" || e.Tag() == "numeric" {
					message = fmt.Sprintf("%s tidak valid", e.Field())
				}
			}
			return errors.New(message)
		}
	}
	return nil
}
