package main

import "fmt"

func main() {
	a, err := NewApp()
	if err != nil {
		panic(fmt.Sprintf("Could not start app: %s", err.Error()))
	}

	a.Run()
}
