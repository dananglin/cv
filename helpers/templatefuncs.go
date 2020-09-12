package helpers

import (
	"fmt"
	"strings"
)

// notLastElement returns true if an element of an array
// or a slice is not the last.
func notLastElement(pos, length int) bool {
	return pos < length-1
}

// join uses strings.Join to join all string elements into
// a single string.
func join(s []string) string {
	return strings.Join(s, " ")
}

// durationToString outputs the employment/education
// duration as a formatted string.
func durationToString(d Duration) string {
	start := fmt.Sprintf("%s, %s", d.Start.Month, d.Start.Year)
	present := strings.ToLower(d.Present)
	end := ""

	if present == "yes" || present == "true" {
		end = "Present"
	} else {
		end = fmt.Sprintf("%s, %s", d.End.Month, d.End.Year)
	}

	return start + " - " + end
}
