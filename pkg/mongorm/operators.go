/*
 * @Author: berryberry
 * @since: 2019-11-08 17:00:06
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
)

func Equal(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}

func NotEqual(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: bson.M{"$ne": value}}
}
