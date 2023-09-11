package main

import "github.com/oktavarium/go-shortener/internal/server"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
