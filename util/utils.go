package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

var r *rand.Rand

const letterStr = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const letterNum = "1234567890"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GBK2UTF8 将GBK转UTF8
func GBK2UTF8(gbkStr string) string {
	return mahonia.NewDecoder("gbk").ConvertString(gbkStr)
}

// GenRandomStrCode 生成随机字符码 默认产出长度为12的随机字符码
func GenRandomStrCode(n int) string {
	if n > 0 {
		sb := strings.Builder{}
		sb.Grow(n)
		for i := 0; i < n; i++ {
			sb.WriteByte(letterStr[r.Int63()%int64(len(letterStr))])
		}
		return sb.String()
	}
	return GenRandomStrCode(12)
}

// GenRandomNumCode 生成随机数字码 默认产出长度为6的随机数字码
func GenRandomNumCode(n int) string {
	if n > 0 {
		sb := strings.Builder{}
		sb.Grow(n)
		for i := 0; i < n; i++ {
			sb.WriteByte(letterNum[r.Int63()%int64(len(letterNum))])
		}
		return sb.String()
	}
	return GenRandomNumCode(6)
}

// VerifyEmail 验证邮箱格式
func VerifyEmail(email string) bool {
	// pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyTelephone 验证手机号格式
func VerifyTelephone(telephone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(telephone)
}

// Md5Encode Md5加密
func Md5Encode(str string) string {
	d := []byte(str)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
