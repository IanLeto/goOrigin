package utils

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


