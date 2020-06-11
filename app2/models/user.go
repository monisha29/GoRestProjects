package models

//what
import "gopkg.in/mgo.v2/bson"

//User struct info
type User struct {
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
	//Id for the user
	ID bson.ObjectId `json:"id" bson:"_id"`
}
