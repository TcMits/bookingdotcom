package model

// Model is an interface that all models must implement
type Model interface {
	CollectionName() string
	SetID(id interface{})
	GetID() interface{}
	IsNew() bool
	IDFieldName() string
}
