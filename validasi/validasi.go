package validasi

import (
	"fmt"

	"github.com/go-playground/validator"
)

func ValidationErrorHandle(err error) string {
	var message string

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("Ada Tabel yang Kosong, Harap Diisi !!!", v.Field())
			case "min":
				message = fmt.Sprintf(" Input Minimal %s character !!!", v.Field(), v.Param())
			case "max":
				message = fmt.Sprintf(" Input Maksimal %s character !!!", v.Field(), v.Param())
			case "lte":
				message = fmt.Sprintf(" Input tidak boleh di bawah %s !!!", v.Field(), v.Param())
			case "gte":
				message = fmt.Sprintf(" Input tidak boleh di atas %s !!!", v.Field(), v.Param())
			case "numeric":
				message = fmt.Sprintf("Input Harus Numeric !!!", v.Field())
			case "url":
				message = fmt.Sprintf(" Input Harus Url !!!", v.Field())
			case "password":
				message = fmt.Sprintf(" Input value must be filled", v.Field())

			}
		}
	}

	return message
}
