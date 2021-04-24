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
	a.Run(":8080")
	fmt.Printf("Application running %s", a.DB.Name())

}
