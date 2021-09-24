package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func EnsureJson(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		logrus.Errorf("数据格式不对 %s", err)
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

func GetFilePath(path string) string {
	return filepath.Join(getRootPath(), path)
}
