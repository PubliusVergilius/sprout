package models

import (
	"encoding/xml"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	XMLName  xml.Name           `xml:"User"`
	ID       primitive.ObjectID `bson:"_id,omitempty" xml:"id,attr"`
	Username string             `bson:"username" xml:"Username"`
	Email    string             `bson:"email" xml:"Email"`
}
