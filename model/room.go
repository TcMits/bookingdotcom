package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	Code              string         `json:"code" bson:"code" fake:"{uuid}"`
	Name              string         `json:"name" bson:"name" fake:"{name}"`
	Description       string         `json:"description" bson:"description" fake:"{sentence:50}"`
	PricePerHour      float64        `json:"pricePerHour" bson:"pricePerHour" fake:"{number:10000,100000000}"`
	PricePerDay       float64        `json:"pricePerDay" bson:"pricePerDay" fake:"{number:10000,100000000}"`
	Images            []Image        `json:"images" bson:"images" fakesize:"1"`
	BedsCount         int            `json:"bedsCount" bson:"bedsCount" fake:"{number:1,10}"`
	MaxAdultGuests    int            `json:"maxAdultGuests" bson:"maxAdultGuests" fake:"{number:1,10}"`
	MaxChildrenGuests int            `json:"maxChildrenGuests" bson:"maxChildrenGuests" fake:"{number:1,10}"`
	ReservedTimes     []ReservedTime `json:"reservedTimes" bson:"reservedTimes" fakesize:"1"`
}

type ReservedTime struct {
	Code        string             `json:"code" bson:"code" fake:"{uuid}"`
	From        primitive.DateTime `json:"from" bson:"from" example:"2021-05-01T00:00:00Z" swaggertype:"primitive,string"`
	To          primitive.DateTime `json:"to" bson:"to" example:"2021-05-01T00:00:00Z" swaggertype:"primitive,string"`
	Name        string             `json:"name" bson:"name" fake:"{name}"`
	Email       string             `json:"email" bson:"email" fake:"{email}"`
	Description string             `json:"description" bson:"description" fake:"{sentence:50}"`
	Phone       string             `json:"phone" bson:"phone" fake:"{phone}"`
	ReceiveTime primitive.DateTime `json:"receiveTime" bson:"receiveTime" example:"2021-05-01T00:00:00Z" swaggertype:"primitive,string"`
	TotalPrice  float64            `json:"totalPrice" bson:"totalPrice" fake:"{number:10000,100000000}"`
}
