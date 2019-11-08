/*
 * @Author: berryberry
 * @since: 2019-11-08 16:58:04
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"context"
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	*mongo.Collection

	transaction *Transaction

	filters bson.D

	Error  error
	Errors []error
}

func (this *Collection) newContext() context.Context {
	if this.transaction != nil {
		return this.transaction.sctx
	}
	return context.Background()
}

func (this *Collection) Where(filter bson.E) *Collection {
	this.filters = append(this.filters, filter)
	return this
}

func (this *Collection) AddError(err error) {
	this.Errors = append(this.Errors)
	this.Error = err
}

func (this *Collection) SelectMany(results interface{}) *Collection {
	if this.Error != nil {
		return this
	}

	cur, err := this.Find(this.newContext(), this.filters)
	if err != nil {
		this.AddError(err)
		return this
	}
	defer cur.Close(this.newContext())

	err = cur.All(this.newContext(), results)
	if err != nil {
		this.AddError(err)
		return this
	}

	elemOfResults := reflect.ValueOf(results).Elem()
	for i := 0; i < elemOfResults.Len(); i++ {
		document := elemOfResults.Index(i).Addr().Interface().(DocumentInterface)
		document.SetCreatedAt(document.GetCreatedAt().Local())
		document.SetUpdatedAt(document.GetUpdatedAt().Local())
		if document.GetDeletedAt() != nil {
			deleteTime := document.GetDeletedAt().Local()
			document.SetDeletedAt(&deleteTime)
		}
	}
	return this
}

func (this *Collection) SelectOne(result DocumentInterface) *Collection {
	if this.Error != nil {
		return this
	}

	var err error
	sr := this.FindOne(this.newContext(), this.filters)
	if err != nil {
		this.AddError(err)
		return this
	}
	sr.Decode(result)
	return this
}

func (this *Collection) CreateOne(document DocumentInterface) *Collection {
	if this.Error != nil {
		return this
	}

	document.SetCreatedAt(time.Now())
	document.SetUpdatedAt(document.GetCreatedAt())

	ir, err := this.InsertOne(this.newContext(), document)
	if err != nil {
		this.AddError(err)
		return this
	}

	// set id
	document.SetID(ir.InsertedID.(primitive.ObjectID))

	// copy createdat, updatedat, deletedat
	document.SetCreatedAt(document.GetCreatedAt().Local())
	document.SetUpdatedAt(document.GetUpdatedAt().Local())
	if document.GetDeletedAt() != nil {
		deleteTime := document.GetDeletedAt().Local()
		document.SetDeletedAt(&deleteTime)
	}

	return this
}

func (this *Collection) ModifyOne(document DocumentInterface) *Collection {
	if this.Error != nil {
		return this
	}

	id := document.GetID()
	if id == primitive.NilObjectID {
		this.AddError(errors.New("Need id, which document.GetID() is primitive.NilObjectID"))
		return this
	}

	var tmpMap map[string]interface{}
	sr := this.FindOne(this.newContext(), bson.M{"_id": id, "deletedat": nil}, options.FindOne().SetProjection(bson.D{{"createdat", 1}}))

	err := sr.Decode(&tmpMap)
	if err != nil {
		this.AddError(err)
		return this
	}

	createdAtTime := tmpMap["createdat"].(primitive.DateTime).Time().Local()

	// update updatedat
	document.SetUpdatedAt(time.Now())
	document.SetCreatedAt(createdAtTime)

	_, err = this.UpdateOne(this.newContext(), bson.M{"_id": id}, bson.M{"$set": document})
	if err != nil {
		this.AddError(err)
		return this
	}

	return this
}

func (this *Collection) RemoveOne(document interface{}) *Collection {
	if this.Error != nil {
		return this
	}

	var err error
	var id primitive.ObjectID

	switch document.(type) {
	case DocumentInterface:
		id = document.(DocumentInterface).GetID()
	case primitive.ObjectID:
		id = document.(primitive.ObjectID)
	default:
		this.AddError(errors.New("type of document is not one of DocumentInterface or primitive.ObjectID"))
		return this
	}

	deletedTime := time.Now()
	_, err = this.UpdateOne(this.newContext(), bson.M{"_id": id}, bson.D{{"$set", bson.D{{"deletedat", &deletedTime}}}})
	if err != nil {
		this.AddError(err)
		return this
	}

	return this
}

func (this *Collection) RemoveMany() *Collection {
	if this.Error != nil {
		return this
	}

	var err error
	deletedTime := time.Now()
	_, err = this.UpdateMany(this.newContext(), this.filters, bson.D{{"$set", bson.D{{"deletedat", &deletedTime}}}})
	if err != nil {
		this.AddError(err)
		return this
	}
	return this
}
