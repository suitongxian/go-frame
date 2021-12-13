package engine

import "strings"

func CreateRouterKey(method string, url string) string {
	var str string
	var build strings.Builder
	build.WriteString(method)
	build.WriteString("-")
	build.WriteString(url)
	str = build.String()
	return str
}
