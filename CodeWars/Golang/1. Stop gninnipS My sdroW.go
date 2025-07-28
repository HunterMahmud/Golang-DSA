package kata
import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	spaceSeparated := strings.Fields(str)
	var result []string
	for i := 0; i < len(spaceSeparated); i++ {
		if len(spaceSeparated[i]) > 4 {
			result = append(result, Reverse(spaceSeparated[i]))
		} else {
			result = append(result, spaceSeparated[i])
		}
	}

	joinedString := strings.Join(result, " ")
	return joinedString
}