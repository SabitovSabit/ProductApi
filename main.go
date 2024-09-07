package main

import (
	"fmt"
	"productapi/dataaccesslayer"
)

func main() {

	dataaccesslayer.StartHttpServer()

	fmt.Println("Hello!")

}
