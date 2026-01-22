package utils

import "github.com/google/uuid"

func IdGenerator() string {
	uuid := uuid.New()
	return uuid.String()
}
