package util

import (
	"fmt"

	"github.com/lithammer/shortuuid/v4"
)

func PrintStruct[T any](a T) string {
	return fmt.Sprintf("%#v", a)
}

func GenerateID() string {
	return shortuuid.New()
}
