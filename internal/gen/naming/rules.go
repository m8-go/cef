package naming

import "strings"

var (
	rules = [...]string{
		"CEF", "ID", "DNS",
	}

	rulesMap = func() (r map[string]string) {
		r = make(map[string]string)
		for _, v := range rules {
			r[strings.ToLower(v)] = v
		}
		return r
	}()
)

func Rule(part string) (string, bool) {
	v, ok := rulesMap[strings.ToLower(part)]
	return v, ok
}
