package utils

import (
	"regexp"
	"strings"
)

func Slugify(input string) string {
	// Map accented characters to their non-accented counterparts
	charMap := map[string]string{
		"á": "a", "à": "a", "ã": "a", "â": "a", "ä": "a",
		"é": "e", "è": "e", "ê": "e", "ë": "e",
		"í": "i", "ì": "i", "î": "i", "ï": "i",
		"ó": "o", "ò": "o", "õ": "o", "ô": "o", "ö": "o",
		"ú": "u", "ù": "u", "û": "u", "ü": "u",
		"ç": "c",
	}

	// Create a regular expression pattern that matches all accented characters
	pattern := regexp.MustCompile("[" + regexp.QuoteMeta(
		"áàãâäéèêëíìîïóòõôöúùûüç") + "]")

	// Replace accented characters with their non-accented counterparts
	output := pattern.ReplaceAllStringFunc(input, func(match string) string {
		return charMap[match]
	})

	re := regexp.MustCompile(`\s+`)
	output = re.ReplaceAllString(output, "-")

	return strings.ToLower(output)
}
