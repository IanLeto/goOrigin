package utils

import "goOrigin/pkg/logging"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorLog(err error) {
	var logger = logging.GetStdLogger()
	if err != nil {
		logger.Errorf("初始化error: %v", err)
	}
}
