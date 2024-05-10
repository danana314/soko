package utilities

import (
	"github.com/lithammer/shortuuid/v4"
)

func NewId() string {
	return shortuuid.New()
}
