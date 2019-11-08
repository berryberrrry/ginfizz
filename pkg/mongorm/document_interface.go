/*
 * @Author: berryberry
 * @since: 2019-11-08 16:58:46
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentInterface interface {
	GetID() primitive.ObjectID
	SetID(primitive.ObjectID)

	SetCreatedAt(time.Time)
	GetCreatedAt() time.Time

	SetUpdatedAt(time.Time)
	GetUpdatedAt() time.Time

	SetDeletedAt(*time.Time)
	GetDeletedAt() *time.Time
}
