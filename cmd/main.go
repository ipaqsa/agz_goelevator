package main

import (
	"go-altitude/pkg/server"
)

const port = ":9111"

func main() {
	err := server.Run(port)
	if err != nil {
		print(err.Error())
		return
	}
}
