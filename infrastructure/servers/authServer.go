package servers

import (
	"encoding/json"
	"fmt"

	"github.com/karolszmaj/gotrack/infrastructure/database/models"
	auth "github.com/karolszmaj/gotrack/infrastructure/service/auth"
	"github.com/karolszmaj/gotrack/infrastructure/service/shared"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//AuthServer Represents Authentication server provider
type AuthServer struct {
	DB *mgo.Database
}

//Login user authentication via email and passwordHash
func (srv *AuthServer) Login(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	cols, err := srv.DB.CollectionNames()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Collection names", cols)

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
}
