package utils

import (
	"github.com/YouthCampProj/douyin/pkg/serializer"
	"mime/multipart"
	"strconv"
)

// StrEncrypt 对传入字符串进行加密
func StrEncrypt(str string, salt string) string {
	// TODO: 实现字符串加密
	return str
}

// StrMatch 对传入的加密字符串进行比对
func StrMatch(str1 string, str2 string, salt string) bool {
	// TODO: 实现加密字符串比对
	return true
}

func UsernameTest(username string) bool {
	return username != ""
	// TODO: 对传入的用户名进行基本的验证(字符|长度等)
}

func PasswordTest(password string) bool {
	return password != ""
	// TODO: 对传入的密码进行基本的验证(字符|长度等)
}

func UserIDTest(userID string) bool {
	return userID != ""
	// TODO: 对传入的用户ID进行基本的验证(字符|长度等)
}

// Str2uint64 将字符串转换为uint64 使用前请确保传入的字符串是合法的
func Str2uint64(str string) uint64 {
	res, _ := strconv.ParseUint(str, 10, 64)
	return res
}

// Str2int64 将字符串转换为int64 使用前请确保传入的字符串是合法的
func Str2int64(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func PublishDataTest(data *multipart.FileHeader) int {
	if data.Size == 0 {
		return serializer.CodePublishFileInvalid
	}
	if data.Filename == "" {
		return serializer.CodePublishFileInvalid
	}
	return serializer.CodeSuccess
	// TODO: 对传入的数据进行基本的验证(文件大小|文件名等)

}
