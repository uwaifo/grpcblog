package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"github.com/uwaifo/grpcblog/global"

	"github.com/uwaifo/grpcblog/proto"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	login, password := in.GetLogin(), in.GetPassword()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()

	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"$or": []bson.M{bson.M{"username": login}, bson.M{"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("Wrong Login Credentials Provided")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("Wrong Login Credentials Provided")
	}

	//return a pointer to the AuthResponse
	return &proto.AuthResponse{Token: user.GetToken()}, nil

}

func main() {

	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("There was an error creating the listener ", err.Error())
	}
	server.Serve(listener)
}
