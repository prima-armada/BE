package validasihp

import (
	"strings"

	"github.com/ttacon/libphonenumber"
)

func ValidateAndFormatPhoneNumber(phoneNumber string) (string, error) {

	countryCode := "ID"
	normalizedNumber := strings.Join(strings.FieldsFunc(phoneNumber, func(r rune) bool {
		return !('0' <= r && r <= '9')
	}), "")

	// Mengecek apakah nomor telepon valid
	number, err := libphonenumber.Parse(normalizedNumber, countryCode)
	if err != nil {
		return "", err
	}

	// Memformat nomor telepon sesuai dengan standar internasional
	formattedNumber := libphonenumber.Format(number, libphonenumber.INTERNATIONAL)
	return formattedNumber, nil
}
