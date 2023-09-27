package base52

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func encode(p []byte) []byte {
	origin := p[:]
	for i := 0; i < len(p)%3; i++ {
		origin = append(origin, 0)
	}
	var result []byte
	for i := 0; i < len(origin); i += 3 {
		result = append(result, alphabet[origin[i]>>2])
		result = append(result, alphabet[(origin[i]<<6)>>4|(origin[i+1]>>4)])
		result = append(result, alphabet[(origin[i+1]<<4)>>4|(origin[i+2]>>6)])
		result = append(result, alphabet[(origin[i+2]<<6)>>6])
	}

	return result
}
