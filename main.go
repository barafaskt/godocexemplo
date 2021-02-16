package main

import (
	"fmt"

	"github.com/barafaskt/exd3/cachorro"
)

// Main principal que mostra o valor na tela em anos caninos
func main() {
	fmt.Println("Anos caninos", cachorro.Idade(5))
}
