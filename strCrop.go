package myutils

import "regexp"

var cropRgx *regexp.Regexp

// StrCrop crops a string like you want in blog previews
func StrCrop(str string, length int) string {
	if len(str) <= length {
		return str
	}

	if cropRgx == nil {
		cropRgx = regexp.MustCompile(`[\s,.-]`)
	}

	var (
		lastStop int
		res      string
	)
	for i, c := range str {
		if cropRgx.MatchString(string(c)) {
			lastStop = i - 1
		}
		if i >= length {
			res = str[:lastStop]
			break
		}
	}

	return res
}
