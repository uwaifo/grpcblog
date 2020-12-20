package main

import (
	"context"
	"testing"

	"github.com/uwaifo/grpcblog/global"

	"github.com/uwaifo/grpcblog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{ID: primitive.NewObjectID(), Email: "jimmychoo@gmail.com", Username: "jimmychoo", Password: string(pw)})
	server := authServer{}
	response, err := server.Login(context.Background(), &proto.LoginRequest{Login: "jimmychoo@gmail.com", Password: "password"})
	if err != nil {
		t.Error("An error was returned : ", err)
	}

	t.Error(response.GetToken())
}
