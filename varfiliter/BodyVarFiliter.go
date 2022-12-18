package varfiliter

import "strings"

func Filiter(from string, to string, text *[]byte) {
	filtered := strings.ReplaceAll(string(*text), "%RECV_EMAIL%", to)
	filtered = strings.ReplaceAll(filtered, "%SEND_EMAIL%", from)
	*text = []byte(filtered)
}
