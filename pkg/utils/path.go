package utils

import (
	"fmt"
	"os"
)

func GetPath(path string) {

}
func getRootPath(path string) string {
	dir, err := os.Getwd()
	NoError(err)
	return fmt.Sprintf("%s/%s", dir, path)
}
