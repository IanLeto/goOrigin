package utils

import (
	"github.com/globalsign/mgo/bson"
)

func ConvBsonNoErr(v interface{}) bson.M {
	res, _ := ConvBson(v)
	return res
}
func ConvBson(v interface{}) (bson.M, error) {
	var doc bson.M
	var err error
	switch value := v.(type) {
	case []byte:
		err = bson.UnmarshalJSON(value, &doc)
		return doc, err
	case string:
		err = bson.UnmarshalJSON([]byte(value), &doc)
		return doc, err
	default:
		data, err := bson.Marshal(v)
		err = bson.Unmarshal(data, &doc)
		return doc, err
	}

}
