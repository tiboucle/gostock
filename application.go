package main

import (
	"gostock/route"
)

func main() {
	err := route.Run(8090)

	if err != nil {
		panic(err)
	}
}
