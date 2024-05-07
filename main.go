package main

import "fmt"

func main() {
	a, err := NewApp()
	if err != nil {
		panic(fmt.Sprintf("Could not create app: %s", err.Error()))
	}

	a.Run()
}
