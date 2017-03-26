package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fluffytime/singularity/app"
	"github.com/fluffytime/singularity/infrastructure/servers"
	"github.com/fluffytime/singularity/infrastructure/service/auth"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
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
