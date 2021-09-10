package utils

import (
	"github.com/cstockton/go-conv"
	"github.com/globalsign/mgo/bson"
)

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

func ConvBson(v interface{}) bson.M {
	var doc bson.M
	data, err := bson.Marshal(v)
	if err != nil {
		return doc
	}
	err = bson.Unmarshal(data, &doc)
	return doc
}
<<<<<<< HEAD

func ConvJsonToBson()  {
}
=======
>>>>>>> d29be502b9084d07f7a09f7d702f05a35f061cca
