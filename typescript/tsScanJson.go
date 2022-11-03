package typescript

import (
	"regexp"
	"strings"
)

type TSTagJson []string

func (t *TSTagJson) parse(tag string) bool {
	*t = []string{}
	re := regexp.MustCompile(`json:\"(.*?)\"`)
	match := re.FindStringSubmatch(tag)
	if len(match) == 0 {
		return false
	}
	s := strings.Split(match[1], ",")
	for _, v := range s {
		*t = append(*t, strings.Trim(v, " "))
	}
	return true
}
