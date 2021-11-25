package utils

func Capitalize(str string) string {
	var upperStr string
	v := []rune(str)
	for i := 0; i < len(v); i++ {
		if i == 0 {
			if v[i] >= 97 && v[i] <= 122 {
				v[i] -= 32
				upperStr += string(v[i])
			}
		} else {
			upperStr += string(v[i])
		}
	}
	return upperStr
}
