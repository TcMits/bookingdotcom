package main

import (
	"context"
	"time"

	"github.com/TcMits/bookingdotcom"
	"github.com/TcMits/bookingdotcom/api"
	"github.com/TcMits/bookingdotcom/core"
	"github.com/TcMits/bookingdotcom/dao"
	"github.com/TcMits/bookingdotcom/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func seed(b bookingdotcom.BookingDotCom) {
  // db, _ := b.DB()
  // db.Collection(model.CollectionNameStays).Drop(context.Background())
	d := dao.NewDao(b)

	count, err := d.Count(context.Background(), &dao.DaoCountConfig{
		CollectionName: model.CollectionNameStays,
		Filter:         bson.D{},
	})

	if err != nil {
		panic(err)
	}

	if count != 0 {
		return
	}

	for i := 0; i < 100; i++ {
		stay := model.Stay{}
		gofakeit.Struct(&stay)
		stay.ID = primitive.NilObjectID

		for i := range stay.Rooms {
			for iv := range stay.Rooms[i].ReservedTimes {
				stay.Rooms[i].ReservedTimes[iv].From = primitive.NewDateTimeFromTime(gofakeit.DateRange(time.Now(), time.Now().Add(time.Hour*24*365)))
				stay.Rooms[i].ReservedTimes[iv].To = primitive.NewDateTimeFromTime(gofakeit.DateRange(time.Now().Add(time.Hour*24*366), time.Now().Add(time.Hour*24*365*2)))
				stay.Rooms[i].ReservedTimes[iv].ReceiveTime = stay.Rooms[i].ReservedTimes[iv].From
			}
		}

		err := d.Mutate(context.Background(), &dao.DaoMutateConfig{
			Model: &stay,
		})

		if err != nil {
			panic(err)
		}
	}

}

func main() {
	mongoClient, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://root:example@bookingdotcom-mongo:27017"),
	)
	defer mongoClient.Disconnect(context.Background())

	if err != nil {
		panic(err)
	}

	app, err := core.NewBookingDotComBuilder().
		SetMongoClient(mongoClient).
		SetDBName("bookingdotcom").
		Build()

	if err != nil {
		panic(err)
	}

	seed(app)
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	api.NewAPIBuilder(app).Build(server)

	if err := server.Start(":8080"); err != nil {
		panic(err)
	}
}
