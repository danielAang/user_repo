package main

import (
	"fmt"
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_URI"),
	)
	a.MountRoutes()
	fmt.Printf("Application running %s", a.DB.Name())
	a.Run(":8080")
}
