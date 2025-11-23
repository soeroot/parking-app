package main

import (
	"fmt"
	"os"

	"soeroot/parking-app/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdout, "required file path")
		fmt.Println("Usage: parking_app <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	parkingService := service.NewParkingService(filePath)
	parkingService.BatchCommandFromFile()
}
