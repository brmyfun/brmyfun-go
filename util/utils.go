package util

import (
	"github.com/axgle/mahonia"
)

// GBK2UTF8 将GBK转UTF8
func GBK2UTF8(gbkStr string) string {
	return mahonia.NewDecoder("gbk").ConvertString(gbkStr)
}
