package servers

import (
	"fmt"

	"github.com/fluffytime/singularity/infrastructure/database/models"
	auth "github.com/fluffytime/singularity/infrastructure/service/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AuthServer Represents Authentication server provider
type AuthServer struct {
	DB *mgo.Database
}

//Login user authentication via email and passwordHash
func (srv *AuthServer) Login(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	fmt.Println("kupa")
	var usr *dbmodels.User
	q := srv.DB.C("users").Find(bson.M{"email": req.Email})
	err := q.One(&usr)

	if err != nil {
		fmt.Println(err.Error())
		return nil, grpc.Errorf(codes.Unauthenticated, "nie ma to nie ma")
	}

	if &usr != nil {
		return &auth.AuthResponse{Token: "baladsada123131312"}, nil
	}

	return nil, fmt.Errorf("Server error")
}

/*

	usr := dbmodels.User{
		Email:     "karolszmaj@whallalabs.com",
		FirstName: "Karol",
		LastName:  "Szmaj",
		Groups:    nil,
		Id:        bson.NewObjectId(),
	}
	srv.DB.C("users").Insert(usr)
	uj, _ := json.Marshal(usr)
	fmt.Printf("%s", uj)

	return &auth.AuthResponse{Status: shared.ResponseStatus_OK, Success: true, Token: "blablabla"}, nil
	/*
		if req.Email == "karol.szmaj@whallalabs.com" && req.PasswordHash == "kurwa" {
			return &auth.AuthResponse{Token: "asdasdasdas", Success: true, Status: shared.ResponseStatus_Unauthorized}, nil
		}

		return nil, fmt.Errorf("Unauthorized")*/
