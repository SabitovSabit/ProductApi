package main

import (
	"fmt"
	"productapi/databaselayer"
)

func main() {
	fmt.Println("Hello!")

	databaselayer.GetConnection()
}
