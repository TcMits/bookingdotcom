package dao

import (
	"context"

	"github.com/TcMits/bookingdotcom"
	"github.com/TcMits/bookingdotcom/core"
	"github.com/TcMits/bookingdotcom/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DaoMutateConfig struct {
	Model         model.Model
	InsertOptions []*options.InsertOneOptions
	UpdateOptions []*options.UpdateOptions
}

type DaoFindConfig struct {
	CollectionName string
	Destination    any
	Filter         any
	Options        []*options.FindOptions
}

type DaoCountConfig struct {
	CollectionName string
	Filter         any
	Options        []*options.CountOptions
}

type Dao struct {
	bookingDotCom bookingdotcom.BookingDotCom
}

func NewDao(bookingDotCom bookingdotcom.BookingDotCom) *Dao {
	return &Dao{
		bookingDotCom: bookingDotCom,
	}
}

func (d *Dao) Mutate(ctx context.Context, config *DaoMutateConfig) error {
	if config == nil {
		return nil
	}

	event := core.NewDaoEvent(config.Model, core.PreMutateEventAction)

	handler := func(e *core.DaoEvent) error {
		db, err := d.bookingDotCom.DB()
		if err != nil {
			return err
		}

		m := e.Model()

		if m.IsNew() {
			m.SetID(primitive.NewObjectID())
			_, err := db.Collection(m.CollectionName()).InsertOne(ctx, m, config.InsertOptions...)
			if err != nil {
				return err
			}

			return nil
		}

		_, err = db.Collection(m.CollectionName()).UpdateOne(
			ctx, bson.M{m.IDFieldName(): m.GetID()}, bson.M{"$set": m}, config.UpdateOptions...,
		)
		if err != nil {
			return err
		}

		return nil
	}

	if err := d.bookingDotCom.DaoHook().Trigger(event, handler); err != nil {
		return err
	}

	event = core.NewDaoEvent(event.Model(), core.PosMutateEventAction)
	if err := d.bookingDotCom.DaoHook().Trigger(event); err != nil {
		return err
	}

	return nil
}

func (d *Dao) Find(ctx context.Context, config *DaoFindConfig) error {
	if config == nil {
		return nil
	}

	db, err := d.bookingDotCom.DB()
	if err != nil {
		return err
	}

	result, err := db.Collection(config.CollectionName).Find(ctx, config.Filter, config.Options...)
	if err != nil {
		return err
	}

	err = result.All(ctx, config.Destination)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) Count(ctx context.Context, config *DaoCountConfig) (int64, error) {
	if config == nil {
		return 0, nil
	}

	db, err := d.bookingDotCom.DB()
	if err != nil {
		return 0, err
	}

	result, err := db.Collection(config.CollectionName).CountDocuments(ctx, config.Filter, config.Options...)
	if err != nil {
		return 0, err
	}

	return result, nil
}
