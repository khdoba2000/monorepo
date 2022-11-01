package hello

import (
	"fmt"
)

func Greet(audience string) string {
	return fmt.Sprintf("from service one to service two, %s!", audience)
}
