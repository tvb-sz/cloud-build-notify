package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

// GenerateFileSha1 通过文件句柄获取文件40位长度的sha1码
func GenerateFileSha1(fp io.Reader) (string, error) {
	sha1Obj := sha1.New()
	_, err := io.Copy(sha1Obj, fp)
	if nil != err {
		return "", err
	}
	return fmt.Sprintf("%x", sha1Obj.Sum(nil)), nil
}

// GenerateUUID 生成一个V4版本的uuid字符串，生成失败返回空串
func GenerateUUID() (string, error) {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}

// StrLen 返回字符串長度 已“”分割的非存儲長度
func StrLen(str string) int {
	return strings.Count(str, "") - 1
}

// StripJSONLiteral 去除json字面量导致解析json对象的特定字符串
func StripJSONLiteral(json string) string {
	if "" == json {
		return ""
	}
	res := strings.Replace(json, "\\", "", -1)
	res = strings.Replace(res, "{", "", -1)
	return strings.Replace(res, "}", "", -1)
}

// Md5 生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Sha1  hash
func Sha1(b []byte) string {
	d := sha1.Sum(b)
	return hex.EncodeToString(d[0:])
}
