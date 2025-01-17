package logic

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goOrigin/internal/model/entity"
	"io"
	"os"
)

func GetConfig(c *gin.Context, path string) ([]byte, error) {
	var (
		err     error
		res     []byte
		file    *os.File
		content []byte
	)
	switch path {
	case "":
		file, err = os.OpenFile("config/configv2.yaml", os.O_RDONLY, 0644)
		if err != nil {
			logger.Error("打开配置文件失败", zap.String("path", path), zap.Error(err))
			goto ERR
		}
	}

	content, err = io.ReadAll(file)
	err = json.Unmarshal(content, &entity.ConfigEntity{})
	if err != nil {
		logger.Error("解析配置文件失败", zap.String("path", path), zap.Error(err))
		goto ERR
	}
	res = content
	return res, err
ERR:
	return nil, err
}
