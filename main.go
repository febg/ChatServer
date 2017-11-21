package main

import (
	"fmt"
	"log"

	"github.com/febg/ChatServer/api"
)

func main() {
	// Create a simple file server

	control, err := api.NewControl(api.ControlConfig{
		LocalDB: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(control)
}
