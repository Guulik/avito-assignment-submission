package main

import (
	"Avito_trainee_assignment/internal/app"
	"fmt"
)

func main() {
	a, err := app.New()
	if err != nil {
		fmt.Println(err)
	}
	err = a.Run()
	if err != nil {
		fmt.Println(err)
	}
}
