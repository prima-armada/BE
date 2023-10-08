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
				message = fmt.Sprintf("%s Ada Tabel yang Kosong, Harap Diisi !!!", v.Field())
			case "min":
				message = fmt.Sprintf("%s Input Minimal %s character !!!", v.Field(), v.Param())
			case "max":
				message = fmt.Sprintf("%s Input Maksimal %s character !!!", v.Field(), v.Param())
			case "lte":
				message = fmt.Sprintf("%s Input tidak boleh di bawah %s !!!", v.Field(), v.Param())
			case "gte":
				message = fmt.Sprintf("%s Input tidak boleh di atas %s !!!", v.Field(), v.Param())
			case "numeric":
				message = fmt.Sprintf("%s Input Harus Numeric !!!", v.Field())
			case "url":
				message = fmt.Sprintf("%s Input Harus Url !!!", v.Field())
			case "password":
				message = fmt.Sprintf("%s Input value must be filled", v.Field())

			}
		}
	}

	return message
}
