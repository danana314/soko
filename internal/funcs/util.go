package funcs

import "fmt"

func PrintStruct(a ...any) string {
	return fmt.Sprintf("%#v", a)
}
