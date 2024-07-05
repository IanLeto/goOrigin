package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

func EnsureJson(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}
	return nil
}

func IncludeString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getRootPath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, "")
}

// 去重
func Set([]string) []string {
	return []string{}
}

func GetFilePath(path string) string {
	return filepath.Join(Root, path)
}

func StrDefault(str, def string) string {
	if str == "" {
		return def
	}
	return str
}

func QueryInt(c *gin.Context, key string, defaultVal ...int) int {
	strv := c.Query(key)
	if strv != "" {
		intv, err := cast.ToIntE(strv)
		if err != nil {
			logrus.Errorf("cannot convert [%s] to int", strv)
		}
		return intv
	}

	if len(defaultVal) == 0 {
		logrus.Errorf("query param[%s] is necessary", key)
	}

	return defaultVal[0]
}

func ToJson(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	NoError(err)

	return string(data)

}

func SafeFunction(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v", r)
		}
	}()
	fn()
}

// GetLoginUrlOrigin 将url的origin部分提取出来，然后给重定向用
func GetLoginUrlOrigin(url string) string {
	return ""
}
