package base52

import "fmt"

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func encode(p []byte) []byte {
	origin := p[:]
	var suffixCount int
	for len(origin)%3 != 0 {
		origin = append(origin, 0)
		suffixCount++
	}
	fmt.Println(len(origin))
	var result []byte
	for i := 0; i < len(origin); i += 3 {
		result = append(result, alphabet[origin[i]>>2])
		result = append(result, alphabet[(origin[i]<<6)>>2|(origin[i+1]>>4)])
		result = append(result, alphabet[(origin[i+1]<<4)>>2|(origin[i+2]>>6)])
		result = append(result, alphabet[(origin[i+2]<<2)>>2])
	}
	for suffixCount > 0 {
		result = append(result, '=')
		suffixCount--
	}

	return result
}
