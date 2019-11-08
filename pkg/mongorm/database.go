/*
 * @Author: berryberry
 * @since: 2019-11-08 16:54:50
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	*mongo.Database
}

func (this *Database) Table(t interface{}, transaction ...*Transaction) *Collection {
	var collection *mongo.Collection

	if namableCollection, ok := t.(NamableCollection); ok {
		collection = this.Collection(namableCollection.CollectionName())
	} else if collectionName, ok := t.(string); ok {
		collection = this.Collection(collectionName)
	} else {
		panic(fmt.Sprintf("unknown table type: %v", t))
	}

	var tran *Transaction
	switch len(transaction) {
	case 0:
		tran = nil
	case 1:
		tran = transaction[0]
	default:
		panic("prime/pkg/mongoloid/database.go: Table() only need 1 or 2 args")
	}

	return &Collection{
		Collection:  collection,
		transaction: tran,
		filters:     bson.D{{"deletedat", nil}},
	}
}

func (this *Database) CreateCollection(t DocumentInterface) error {
	deleteTime := time.Now()
	t.SetDeletedAt(&deleteTime)

	ir, err := this.Table(t).InsertOne(context.Background(), t)
	if err != nil {
		return err
	}

	objectID := ir.InsertedID.(primitive.ObjectID)
	_, err = this.Table(t).DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
