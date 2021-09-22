package utils

import (
	"github.com/globalsign/mgo/bson"
)


func ConvBson(v interface{}) bson.M {
	var doc bson.M
	data, err := bson.Marshal(v)
	if err != nil {
		return doc
	}
	err = bson.Unmarshal(data, &doc)
	return doc
}
