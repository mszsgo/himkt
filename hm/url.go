package hm

import "strings"

func UrlAddParams(url, k, v string) string {
	if strings.Contains(url, "#") {
		ss := strings.Split(url, "#")
		s2 := ss[1]
		if strings.Contains(s2, "?") {
			s2 += "&"
		} else {
			s2 += "?"
		}
		s2 += k + "=" + v
		return ss[0] + "#" + s2
	}

	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}
	url += k + "=" + v
	return url
}
