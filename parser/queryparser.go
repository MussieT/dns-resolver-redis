package parser

import (
	"fmt"
	"strings"
)

// Explain - further (bytes, string builder and others..)
func removeNonAlphaNumeric(s []byte) []byte {
    n := 0
    for _, b := range s {
        if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') ||
            ('0' <= b && b <= '9') ||
            b == '-' || b == '.' {
            s[n] = b
            n++
        }
    }
    return s[:n]
}

func removePrefixIfExists(s string) string {

	prefixes := []string{"http://", "https://", "www."}

	after := s

	for _, prefix := range prefixes {
		if strings.Contains(s, prefix) {
			after = strings.TrimPrefix(s, prefix)
		}
	}

	return after
}

func SanitizeQuery(domain string) string {

	trimmedDomain := removePrefixIfExists((domain))

	trimmedDomain = string(removeNonAlphaNumeric([]byte(trimmedDomain)))

	if len(domain) > 25 {
		fmt.Println("Length reached max")
	}

	return trimmedDomain

}
