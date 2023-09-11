package main

import "github.com/oktavarium/go-shortener/internal/shortener"

func main() {
	if err := shortener.Run(); err != nil {
		panic(err)
	}
}
