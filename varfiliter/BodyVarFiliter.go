package varfiliter

import (
	"SimpleDecrypt/datastruct"
	"strings"
)

func VarFiliter(from string, to string, text *[]byte) {
	filtered := strings.ReplaceAll(string(*text), "%RECV_EMAIL%", to)
	filtered = strings.ReplaceAll(filtered, "%SEND_EMAIL%", from)
	*text = []byte(filtered)
}

func FormatFiliter(from *string) {
	*from = strings.ReplaceAll(*from, "\n", "")
	*from = strings.ReplaceAll(*from, "\r", "")
	*from = strings.ReplaceAll(*from, "\t", "")
}

func GroupFiliter(sendPackages *[]datastruct.EmailStructs) {
	var filiteredPackage = *sendPackages
	for _, it := range filiteredPackage {
		body := it.Text
		VarFiliter(it.From, it.To[0], &body)
	}
	*sendPackages = filiteredPackage
}
