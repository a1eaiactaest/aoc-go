package algo

import "strings"

func SplitOn(in string, cutset []string) []string {
	parts := strings.Split(in, cutset[0])
	cutset = cutset[1:] 
	var finished bool

	for !finished && len(cutset) > 0 {
		divider := cutset[0]
		var newParts []string
		for _, oldPart := range parts {
			newParts = append(newParts, strings.Split(oldPart, divider)...)
		}
		parts = newParts
	}
	return parts
}
