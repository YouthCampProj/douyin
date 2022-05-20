package utils

import (
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/YouthCampProj/douyin/pkg/serializer"
	"golang.org/x/crypto/bcrypt"
)

// StrEncrypt 对传入字符串进行加密
func StrEncrypt(str string, salt string) string {
	// TODO: 实现字符串加密
	str = str + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// StrMatch 对传入的加密字符串进行比对,str2为明文
func StrMatch(str1 string, str2 string, salt string) bool {
	// TODO: 实现加密字符串比对
	str2 = str2 + salt
	err := bcrypt.CompareHashAndPassword([]byte(str1), []byte(str2))
	return err == nil
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

func VideoIDTest(videoID string) bool {
	return videoID != ""
	// TODO: 对传入的视频ID进行基本的验证(字符|长度等)
}

func FavoriteActionTypeTest(actionType string) bool {
	switch actionType {
	case "1":
		return true
	case "2":
		return true
	default:
		return false
	}
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

// Str2int32 将字符串转换为int32 使用前请确保传入的字符串是合法的
func Str2int32(str string) int32 {
	res, _ := strconv.ParseInt(str, 10, 32)
	return int32(res)
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

func GetExtensionName(fileName string) string {
	return strings.ToLower(filepath.Ext(fileName))
}
