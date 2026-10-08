package greet

import (
	"fmt"
	"strings"
)

func Greet(name, salutation string) string {
	if strings.TrimSpace(name) == "" {
		return salutation
	}
	return fmt.Sprintf("%s, %s", salutation, name)
}
