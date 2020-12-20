module github.com/uwaifo/grpcblog

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/crypto v0.0.0-20201217014255-9d1352758620
	google.golang.org/grpc v1.34.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/uwaifo/grpcblog/backend/AuthService => ./backend/AuthService
