package utils

import "github.com/cstockton/go-conv"

func ConvOrDefaultString(from interface{}, defaults string) string {
	v, err := conv.String(from)
	if err != nil {
		v = defaults
	}
	return v
}
