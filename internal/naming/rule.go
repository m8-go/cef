package naming

import "strings"

var (
	rules = [...]string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC", "MB",
		"QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO", "TLS",
		"TTL", "UI", "UID", "URI", "URL", "UTF8", "UUID", "VM", "XML", "XMPP",
		"XSRF", "XSS", "SMS", "CDN", "TCP", "UDP", "DC", "PFS", "P2P",
		"SHA256", "SHA1", "MD5", "SRP", "2FA", "OAuth", "OAuth2",

		"PNG", "JPG", "GIF", "MP4", "WEBP",
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
