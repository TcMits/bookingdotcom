package core

import (
	"context"
	"fmt"

	"github.com/pocketbase/pocketbase/tools/hook"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingDotComBuilder struct {
	mongoClient *mongo.Client
	dbName      string
	daoHook     *hook.Hook[*DaoEvent]
}

func NewBookingDotComBuilder() *BookingDotComBuilder {
	return &BookingDotComBuilder{
		daoHook: new(hook.Hook[*DaoEvent]),
	}
}

func (b *BookingDotComBuilder) SetMongoClient(mongoClient *mongo.Client) *BookingDotComBuilder {
	b.mongoClient = mongoClient
	return b
}

func (b *BookingDotComBuilder) SetDBName(dbName string) *BookingDotComBuilder {
	b.dbName = dbName
	return b
}

func (b *BookingDotComBuilder) RegisterDaoHookHandler(daoHook hook.Handler[*DaoEvent]) *BookingDotComBuilder {
	b.daoHook.Add(daoHook)
	return b
}

func (b *BookingDotComBuilder) Build() (*BookingDotCom, error) {
	if b.mongoClient == nil {
		return nil, fmt.Errorf("mongo client is required")
	}

	return &BookingDotCom{
		mongoClient: b.mongoClient,
		dbName:      b.dbName,
		daoHook:     b.daoHook,
	}, nil
}

type BookingDotCom struct {
	mongoClient *mongo.Client
	dbName      string

	daoHook *hook.Hook[*DaoEvent]
}

// getMongoClient returns a mongo client
func (b *BookingDotCom) getMongoClient(ctx context.Context) (*mongo.Client, error) {
	if err := b.mongoClient.Ping(ctx, nil); err == nil {
		return b.mongoClient, nil
	}

	if err := b.mongoClient.Connect(ctx); err != nil {
		return nil, err
	}

	return b.mongoClient, nil
}

// DB returns the mongo database
func (b *BookingDotCom) DB() (*mongo.Database, error) {
	client, err := b.getMongoClient(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database(b.dbName), nil
}

// DaoHook returns the dao hook
func (b *BookingDotCom) DaoHook() *hook.Hook[*DaoEvent] {
	return b.daoHook
}
