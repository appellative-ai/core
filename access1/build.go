package access1

import (
	"net/url"
	"strings"
)

// BuildQuery - build a query string with escaping
func BuildQuery(query any) string {
	if query == nil {
		return ""
	}
	if s, ok := query.(string); ok {
		return BuildValues(s).Encode()
	}
	if v, ok := query.(url.Values); ok {
		return v.Encode()
	}
	return ""
}

// BuildValues - build an url.Values type
func BuildValues(query string) url.Values {
	values := make(url.Values)
	if query == "" {
		return values
	}
	args := strings.Split(query, "&")
	if len(args) == 0 {
		return values
	}
	value := ""
	for _, v := range args {
		pair := strings.Split(v, "=")
		if len(pair) == 1 || pair[1] == "" {
			value = "invalid"
		} else {
			value = pair[1]
		}
		values.Add(pair[0], value)
	}
	return values
}
