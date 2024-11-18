package hello

import "fmt"

func Greet(audience string) string {
	return fmt.Sprintf("Привет, %s!", audience)
}
