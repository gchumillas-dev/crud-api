package str

import "unicode"

// Ucfirst makes a string's first character uppercase.
//
// https://www.php2golang.com/method/function.ucfirst.html
//
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}
