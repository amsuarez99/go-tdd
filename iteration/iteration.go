package iteration

func Repeat(char string, times int) string {
	var out string
	for range times {
		out += char
	}

	return out
}
