package utils

import "goOrigin/logging"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckNoError(err error) {
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
