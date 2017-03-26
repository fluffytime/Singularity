package main

import (
	"fmt"
	"log"
	"net"
	"github.com/karolszmaj/gotrack/infrastructure/service/auth"
	"gopkg.in/mgo.v2"
	"github.com/karolszmaj/gotrack/app"
	"github.com/karolszmaj/gotrack/infrastructure/servers"
	"google.golang.org/grpc"
)

var dbs *mgo.Database

func logSQL(a string, args ...interface{}) {
	fmt.Println(a, args)
}

func main() {
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration %s", err))
	}

	//init database
	dbs, err := mgo.Dial(app.Config.DSN)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	auth.RegisterAuthServer(srv, &servers.AuthServer{
		DB: dbs.DB("gotrack"),
	})
	srv.Serve(lis)
}