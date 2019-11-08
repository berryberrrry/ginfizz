/*
 * @Author: berryberry
 * @since: 2019-11-08 16:54:08
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package mongorm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
	db     *Database
}

func New(username, password, host, dbname string) (*Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", username, password, host, dbname)
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = mongoClient.Connect(nil)
	if err != nil {
		return nil, err
	}

	db := Database{mongoClient.Database(dbname)}
	client := Client{mongoClient, &db}
	return &client, nil
}

func (this *Client) DB() *Database {
	return this.db
}

func (this *Client) ExecuteTransaction(fn func(*Transaction) error) error {
	transaction := new(Transaction)
	return this.client.UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		transaction.sctx = sessionContext
		return fn(transaction)
	})
}
