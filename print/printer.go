package print

import "fmt"

type Printer interface {
	Print(message string)
}

type StandardOutput struct{}

func (so StandardOutput) Print(message string) {
	fmt.Println(message)
}
