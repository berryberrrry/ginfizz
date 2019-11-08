/*
 * @Author: berryberry
 * @since: 2019-11-08 16:57:05
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction struct {
	sctx mongo.SessionContext
}

type TransactionFunc func(*Transaction) error

func NewMongoDBTransaction() *Transaction {
	transaction := new(Transaction)
	return transaction
}

func (this *Transaction) Context() mongo.SessionContext {
	return this.sctx
}

func (this *Transaction) Begin() error {
	return this.sctx.StartTransaction()
}

func (this *Transaction) Commit() error {
	return this.sctx.CommitTransaction(this.sctx)
}

func (this *Transaction) Rollback() error {
	return this.sctx.AbortTransaction(this.sctx)
}

func (this *Transaction) RecoverAndRollback() {
	if r := recover(); r != nil {
		this.Rollback()
	}
}
