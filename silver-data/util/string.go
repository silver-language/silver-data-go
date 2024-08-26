package util

import "regexp"

type Submatch struct {
	String string
	Start  int
	End    int
}

func FindSubmatches(string string, regex regexp.Regexp) []Submatch {
	result := []Submatch{}

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)

	for i := 0; i < len(stringSubmatchIndex); i = i + 2 {
		result = append(result,
			Submatch{
				String: string[stringSubmatchIndex[i]:stringSubmatchIndex[i+1]],
				Start:  stringSubmatchIndex[i],
				End:    stringSubmatchIndex[i+1],
			},
		)
	}

	return result
}
