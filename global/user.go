package global

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NillUser is the nil value of type User
var NilUser User

//User is the default user struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson: "password"`
}

func (u User) GetToken() string {
	byteSlice, _ := json.Marshal(u)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"data": string(byteSlice)})

	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString

}
