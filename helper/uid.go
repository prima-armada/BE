package helper

import "github.com/google/uuid"

func convertUid() string {
	return uuid.New().String()
}
