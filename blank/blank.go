package blank

import "unicode"

func Remove(str string)  string {
	var out []rune
	for _, r := range str {
		if !unicode.IsSpace(r) {
			out = append(out, r)
		}
	}
	return string(out)
}

func Is(str string) bool {
	if Remove(str) == "" {
		return true
	}
	return false
}

func Has(slice []string) bool  {
	if len(slice) <= 0 {
		return true
	}

	for _, s := range slice {
		if Is(s) {
			return true
		}
	}
	return false
}