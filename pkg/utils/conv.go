package utils

import "github.com/cstockton/go-conv"

func ConvOrDefaultString(from interface{}, defaults string) string {
	v, err := conv.String(from)
	if err != nil {
		v = defaults
	}
	return v
}

func ConvOrDefaultInt(from interface{}, defaults int) int {
	v, err := conv.Int(from)
	if err != nil {
		v = defaults
	}
	return v
}


