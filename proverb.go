package proverb

import (
	"fmt"
)

func Proverb(rhyme []string) []string {

	var output []string

	for count, str := range rhyme {
		var verb string
		if count > len(rhyme)-2 {
			verb = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			verb = fmt.Sprintf("For want of a %s the %s was lost.", str, rhyme[count+1])
		}
		output = append(output, verb)
	}
	return output
}
