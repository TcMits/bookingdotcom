package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var _ Model = (*Stay)(nil)

const CollectionNameStays = "stays"

type Stay struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name" fake:"{name}"`
	Rooms       []Room             `json:"rooms" bson:"rooms" fakesize:"1"`
	Description string             `json:"description" bson:"description" fake:"{sentence:50}"`
	Images      []Image            `json:"images" bson:"images" fakesize:"1"`
	StayType    string             `json:"stayType" bson:"stayType" fake:"{randomstring:[hotel,apartment,hostel]}"`
	Address     `json:",inline" bson:"inline"`
}

func (s *Stay) CollectionName() string {
	return CollectionNameStays
}

func (s *Stay) SetID(id interface{}) {
	s.ID = id.(primitive.ObjectID)
}

func (s *Stay) GetID() interface{} {
	return s.ID
}

func (s *Stay) IsNew() bool {
	return s.ID.IsZero()
}

func (s *Stay) IDFieldName() string {
	return "_id"
}
