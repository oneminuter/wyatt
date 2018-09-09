package util

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"unicode"
	"wyatt/api/constant"
)

const (
	num = "0123456789"
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

/*获取定长随机字符
l 长度,要生成的字符串长度
t 类型,number 纯数字类型, string 数字字母组合类型
*/

func GetRandomString(length int, rtype string) string {
	res := ""
	if rtype == "number" {
		res = num
	} else if rtype == "string" {
		res = str
	} else {
		return ""
	}

	bytes := []byte(res)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
当前纳秒 + 10位随机字符串 的MD5值
*/
func GetUUID() string {
	return MD5(strconv.FormatInt(time.Now().UnixNano(), 10) + GetRandomString(10, constant.STRING))
}

/*
	判断字符串是否包含中文及中文符号
*/
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

/*
验证手机号是否合法
*/
func IsPhoneNumber(str string) bool {
	reg := regexp.MustCompile(`^(1[358][0-9]|14[57]|17[0678]|197)\d{8}$`)
	return reg.MatchString(str)
}

/*
验证是否为邮箱
条件：
1. @之前必须有内容且只能是字母（大小写）、数字、下划线(_)、减号（-）、点（.）
2. @和最后一个点（.）之间必须有内容且只能是字母（大小写）、数字、点（.）、减号（-），且两个点不能挨着
3. 最后一个点（.）之后必须有内容且内容只能是字母（大小写）、数字且长度为大于等于2个字节，小于等于6个字节
*/
func IsEmail(str string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`)
	return reg.MatchString(str)
}
