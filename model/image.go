package model

type Image struct {
  FileName string `json:"fileName" bson:"fileName" fake:"https://picsum.photos/200/300"`
  AltText  string `json:"altText" bson:"altText" fake:"{sentence:10}"`
}
